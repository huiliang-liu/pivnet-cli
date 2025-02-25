package commands

import (
	"fmt"
	"github.com/pivotal-cf/go-pivnet/v2/logshim"
	"github.com/pivotal-cf/pivnet-cli/filter"
	"github.com/pivotal-cf/pivnet-cli/rc/filesystem"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/go-pivnet/v2/logger"
	"github.com/pivotal-cf/go-pivnet/v2/sha256sum"
	"github.com/pivotal-cf/pivnet-cli/auth"
	"github.com/pivotal-cf/pivnet-cli/errorhandler"
	"github.com/pivotal-cf/pivnet-cli/gp"
	"github.com/pivotal-cf/pivnet-cli/printer"
	"github.com/pivotal-cf/pivnet-cli/rc"
	"github.com/pivotal-cf/pivnet-cli/version"
	"github.com/robdimsdale/sanitizer"
)

//go:generate counterfeiter . Authenticator
type Authenticator interface {
	AuthenticateClient(client auth.AuthClient) error
}

type Filterer interface {
	ReleasesByVersion(releases []pivnet.Release, version string) ([]pivnet.Release, error)
	ProductFileKeysByGlobs(productFiles []pivnet.ProductFile, globs []string) ([]pivnet.ProductFile, error)
}

//go:generate counterfeiter . RCHandler
type RCHandler interface {
	SaveProfile(profileName string, apiToken string, host string, accessToken string, accessTokenExpiry int64) error
	ProfileForName(profileName string) (*rc.PivnetProfile, error)
	RemoveProfileWithName(profileName string) error
}

var (
	OutputWriter io.Writer
	LogWriter    io.Writer

	Filter           Filterer
	ErrorHandler     errorhandler.ErrorHandler
	Printer          printer.Printer
	RC               RCHandler
	Auth             Authenticator
	Sha256FileSummer sha256sum.FileSummer
)

type PivnetCommand struct {
	VersionFunc func() `short:"v" long:"version" description:"Print the version of this CLI and exit"`

	Format  string `long:"format" description:"Format to print as" default:"table" choice:"table" choice:"json" choice:"yaml"`
	Verbose bool   `long:"verbose" description:"Display verbose output"`

	ProfileName       string `long:"profile" description:"Name of profile" default:"default"`
	ConfigFile        string `long:"config" description:"Path to config file"`
	SkipSSLValidation bool   `long:"skip-ssl-validation" description:"Skip verification of the API endpoint. Not recommended!"`

	Login  LoginCommand  `command:"login" alias:"l" description:"Log in to Pivotal Network."`
	Logout LogoutCommand `command:"logout" description:"Log out from Pivotal Network."`

	Help    HelpCommand    `command:"help" alias:"h" description:"Print this help message"`
	Version VersionCommand `command:"version" alias:"v" description:"Print the version of this CLI and exit"`

	Curl CurlCommand `command:"curl" alias:"c" description:"Curl an endpoint"`

	ReleaseTypes ReleaseTypesCommand `command:"release-types" alias:"rts" description:"List release types"`

	PivnetVersions PivnetVersionsCommand `command:"pivnet-versions" alias:"pv" description:"List Pivnet product versions"`

	EULAs      EULAsCommand      `command:"eulas" alias:"es" description:"List EULAs"`
	EULA       EULACommand       `command:"eula" alias:"e" description:"Show EULA"`
	AcceptEULA AcceptEULACommand `command:"accept-eula" alias:"ae" description:"Accept EULA (Available for pivots only)"`

	Products ProductsCommand `command:"products" alias:"ps" description:"List products"`
	Product  ProductCommand  `command:"product" alias:"p" description:"Show product"`
	ProductSlugs ProductSlugsCommand `command:"product-slugs" alias:"psl" description:"Show slugs associated to a product"`

	ProductFiles      ProductFilesCommand      `command:"product-files" alias:"pfs" description:"List product files"`
	ProductFile       ProductFileCommand       `command:"product-file" alias:"pf" description:"Show product file"`
	CreateProductFile CreateProductFileCommand `command:"create-product-file" alias:"cpf" description:"Create product file"`
	UpdateProductFile UpdateProductFileCommand `command:"update-product-file" alias:"upf" description:"Update product file"`
	AddProductFile    AddProductFileCommand    `command:"add-product-file" alias:"apf" description:"Add product file to release"`
	RemoveProductFile RemoveProductFileCommand `command:"remove-product-file" alias:"rpf" description:"Remove product file from release"`
	DeleteProductFile DeleteProductFileCommand `command:"delete-product-file" alias:"dpf" description:"Delete product file"`

	DownloadProductFiles DownloadProductFilesCommand `command:"download-product-files" alias:"dlpf" description:"Download product files"`

	FileGroups                 FileGroupsCommand                 `command:"file-groups" alias:"fgs" description:"List file groups"`
	FileGroup                  FileGroupCommand                  `command:"file-group" alias:"fg" description:"Show file group"`
	CreateFileGroup            CreateFileGroupCommand            `command:"create-file-group" alias:"cfg" description:"Create file group"`
	UpdateFileGroup            UpdateFileGroupCommand            `command:"update-file-group" alias:"ufg" description:"Update file group"`
	DeleteFileGroup            DeleteFileGroupCommand            `command:"delete-file-group" alias:"dfg" description:"Delete file group"`
	AddFileGroupToRelease      AddFileGroupToReleaseCommand      `command:"add-file-group" alias:"afg" description:"Add file group to release"`
	RemoveFileGroupFromRelease RemoveFileGroupFromReleaseCommand `command:"remove-file-group" alias:"rfg" description:"Remove file group from release"`

	ImageReferences                 ImageReferencesCommand                 `command:"image-references" alias:"irs" description:"List image references"`
	ImageReference                  ImageReferenceCommand                  `command:"image-reference" alias:"ir" description:"Show image reference"`
	CreateImageReference            CreateImageReferenceCommand            `command:"create-image-reference" alias:"cir" description:"Create a container image reference"`
	DeleteImageReference            DeleteImageReferenceCommand            `command:"delete-image-reference" alias:"dir" description:"Delete a container image reference"`
	AddImageReferenceToRelease      AddImageReferenceToReleaseCommand      `command:"add-image-reference" alias:"air" description:"Add image reference to release"`
	RemoveImageReferenceFromRelease RemoveImageReferenceFromReleaseCommand `command:"remove-image-reference" alias:"rir" description:"Remove image reference from release"`
	UpdateImageReference            UpdateImageReferenceCommand            `command:"update-image-reference" alias:"uir" description:"Update a container image reference"`

	Releases      ReleasesCommand      `command:"releases" alias:"rs" description:"List releases"`
	Release       ReleaseCommand       `command:"release" alias:"r" description:"Show release"`
	CreateRelease CreateReleaseCommand `command:"create-release" alias:"cr" description:"Create release"`
	DeleteRelease DeleteReleaseCommand `command:"delete-release" alias:"dr" description:"Delete release"`
	UpdateRelease UpdateReleaseCommand `command:"update-release" alias:"ur" description:"Update release"`

	UserGroups      UserGroupsCommand      `command:"user-groups" alias:"ugs" description:"List user groups"`
	UserGroup       UserGroupCommand       `command:"user-group" alias:"ug" description:"Show user group"`
	AddUserGroup    AddUserGroupCommand    `command:"add-user-group" alias:"aug" description:"Add user group to release"`
	RemoveUserGroup RemoveUserGroupCommand `command:"remove-user-group" alias:"rug" description:"Remove user group from release"`
	CreateUserGroup CreateUserGroupCommand `command:"create-user-group" alias:"cug" description:"Create user group"`
	UpdateUserGroup UpdateUserGroupCommand `command:"update-user-group" alias:"uug" description:"Update user group"`
	DeleteUserGroup DeleteUserGroupCommand `command:"delete-user-group" alias:"dug" description:"Delete user group"`

	AddUserGroupMember    AddUserGroupMemberCommand    `command:"add-user-group-member" alias:"augm" description:"Add user group member to group"`
	RemoveUserGroupMember RemoveUserGroupMemberCommand `command:"remove-user-group-member" alias:"rugm" description:"Remove user group member from group"`

	CompanyGroups CompanyGroupsCommand `command:"company-groups" alias:"cgs" description:"List managed company groups"`
	CompanyGroup CompanyGroupCommand `command:"company-group" alias:"cg" description:"Show company group"`
	AddCompanyGroupMember AddCompanyGroupMemberCommand `command:"company-group-add-member" alias:"cgam" description:"Add a member to a company group"`
	RemoveCompanyGroupMember RemoveCompanyGroupMemberCommand `command:"company-group-remove-member" alias:"cgrm" description:"Remove a member to a company group"`

	ReleaseDependencies     ReleaseDependenciesCommand     `command:"release-dependencies" alias:"rds" description:"List release dependencies"`
	AddReleaseDependency    AddReleaseDependencyCommand    `command:"add-release-dependency" alias:"ard" description:"Add release dependency"`
	RemoveReleaseDependency RemoveReleaseDependencyCommand `command:"remove-release-dependency" alias:"rrd" description:"Remove release dependency"`

	DependencySpecifiers      DependencySpecifiersCommand      `command:"dependency-specifiers" alias:"dss" description:"List dependency specifiers"`
	DependencySpecifier       DependencySpecifierCommand       `command:"dependency-specifier" alias:"ds" description:"Get dependency specifier"`
	CreateDependencySpecifier CreateDependencySpecifierCommand `command:"create-dependency-specifier" alias:"cds" description:"Create dependency specifier"`
	DeleteDependencySpecifier DeleteDependencySpecifierCommand `command:"delete-dependency-specifier" alias:"dds" description:"Delete dependency specifier"`

	ReleaseUpgradePaths      ReleaseUpgradePathsCommand      `command:"release-upgrade-paths" alias:"rups" description:"List release upgrade paths"`
	AddReleaseUpgradePath    AddReleaseUpgradePathCommand    `command:"add-release-upgrade-path" alias:"arup" description:"Add release upgrade path"`
	RemoveReleaseUpgradePath RemoveReleaseUpgradePathCommand `command:"remove-release-upgrade-path" alias:"rrup" description:"Remove release upgrade path"`

	Logger    logger.Logger
	userAgent string
	profile   *rc.PivnetProfile
}

var Pivnet PivnetCommand

func init() {
	Pivnet.VersionFunc = func() {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	if Pivnet.ConfigFile == "" {
		userHomeDir, err := userHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		Pivnet.ConfigFile = filepath.Join(userHomeDir, ".pivnetrc")
	}
}

func NewPivnetClient() *gp.Client {
	var refreshToken string
	var host string

	if Pivnet.profile != nil {
		refreshToken = Pivnet.profile.APIToken
		host = Pivnet.profile.Host
	}

	accessTokenService := CreateAccessTokenService(RC, Pivnet.ProfileName, refreshToken, host, Pivnet.SkipSSLValidation)
	return NewPivnetClientWithToken(accessTokenService, host)
}

func NewPivnetClientWithToken(tokenService gp.AccessTokenService, host string) *gp.Client {
	config := pivnet.ClientConfig{
		Host:              host,
		UserAgent:         Pivnet.userAgent,
		SkipSSLValidation: Pivnet.SkipSSLValidation,
	}
	return gp.NewClient(
		tokenService,
		config,
		Pivnet.Logger,
	)
}

var Init = func(profileRequired bool) error {
	if OutputWriter == nil {
		OutputWriter = os.Stdout
	}

	if LogWriter == nil {
		switch Pivnet.Format {
		case printer.PrintAsJSON, printer.PrintAsYAML:
			LogWriter = os.Stderr
			break
		default:
			LogWriter = os.Stdout
		}
	}

	if ErrorHandler == nil {
		ErrorHandler = errorhandler.NewErrorHandler(Pivnet.Format, OutputWriter, LogWriter)
	}

	if Auth == nil {
		Auth = auth.NewAuthenticator(ErrorHandler)
	}

	if Printer == nil {
		Printer = printer.NewPrinter(OutputWriter)
	}

	if RC == nil {
		rcFileReadWriter := filesystem.NewPivnetRCReadWriter(Pivnet.ConfigFile)
		RC = rc.NewRCHandler(rcFileReadWriter)
	}

	profile, err := RC.ProfileForName(Pivnet.ProfileName)
	if err != nil {
		return ErrorHandler.HandleError(err)
	}

	if profileRequired {
		if profile == nil {
			err := fmt.Errorf("Please login first")
			return ErrorHandler.HandleError(err)
		} else {
			err := profile.Validate()
			if err != nil {
				e := fmt.Errorf("Saved profile is invalid (%s). Please login again", err.Error())
				return ErrorHandler.HandleError(e)
			}
		}
	}

	if profile != nil {
		Pivnet.profile = profile

		sanitizeWriters(profile.APIToken)
	}

	infoLogger := log.New(LogWriter, "", log.LstdFlags)
	debugLogger := log.New(LogWriter, "", log.LstdFlags)

	Pivnet.userAgent = fmt.Sprintf(
		"pivnet-cli/%s",
		version.Version,
	)

	Pivnet.Logger = logshim.NewLogShim(infoLogger, debugLogger, Pivnet.Verbose)

	if Filter == nil {
		Filter = filter.NewFilter(
			Pivnet.Logger,
		)
	}

	return nil
}

func userHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home != "" {
		return home, nil
	}

	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		if home != "" {
			return home, nil
		}

		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home != "" {
			return home, nil
		}
	}

	return "", fmt.Errorf("could not detect home directory for .pivnetrc")
}

func sanitizeWriters(apiToken string) {
	sanitized := map[string]string{
		apiToken: "*** redacted api token ***",
	}

	OutputWriter = sanitizer.NewSanitizer(sanitized, OutputWriter)
	LogWriter = sanitizer.NewSanitizer(sanitized, LogWriter)
}
