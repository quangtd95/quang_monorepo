workspace(name = "quang_monorepo")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
    ],
    sha256 = "513c12397db1bc9aa46dd62f02dd94b49a9b5d17444d49b5a04c5a89f3053c1c",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Download the rules_docker repository at release v0.13.0
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "df13123c44b4a4ff2c2f337b906763879d94871d16411bf82dcfeba892b58607",
    strip_prefix = "rules_docker-0.13.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.13.0/rules_docker-v0.13.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_labstack_echo_v4",
    importpath = "github.com/labstack/echo/v4",
    sum = "h1:h8XP66UfB3tUm+L3QPw7tmwAu3pJaA/nyfHPCcz46ic=",
    version = "v4.1.14",
)

go_repository(
    name = "com_github_labstack_gommon",
    importpath = "github.com/labstack/gommon",
    sum = "h1:JEeO0bvc78PKdyHxloTKiF8BD5iGrH8T6MSeGvSgob0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:snbPLB8fVfU9iwbbo30TPtbLRzwWu6aJS6Xh4eaaviA=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:FxPOTFNqGkuDUGi3H/qkUbQO4ZiBa2brKq5r0l8TGeM=",
    version = "v0.0.11",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:RZqt0yGBsps8NGvLSGW804QQqCUYYLsaOjTVHy1Ocw4=",
    version = "v1.1.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version = "v2.2.2",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:sKJQZMuxjOAR/Uo2LBfU90onWEf1dF4C+0hPJCc9Mpc=",
    version = "v0.0.0-20191227163750-53104e6ec876",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:efeOvDhwQ29Dj3SdAV/MJf8oukgn+8D8WgaCaRMchF8=",
    version = "v0.0.0-20191209160850-c0dbc17a3553",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:JA8d3MPx/IToSyXZG/RhwYEtfrKO1Fxrqe8KrkiLXKM=",
    version = "v0.0.0-20191228213918-04cbcbbfeed8",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
    version = "v0.3.2",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:FDhOuMEY4JVRztM/gsbk+IKUQ8kj74bxZrgw87eMMVc=",
    version = "v0.0.0-20180917221912-90fa682c2a6e",
)
