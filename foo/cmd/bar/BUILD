load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# First we load the go_image and container_push rules
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "quang_monorepo/foo/cmd/bar",
    visibility = ["//visibility:private"],
    deps = ["@com_github_labstack_echo_v4//:go_default_library",
            "//foo/pkg:go_default_library"],
)

go_binary(
    name = "bar",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "docker",
    binary = ":bar",
)
