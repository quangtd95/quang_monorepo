load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["utils.go"],
    importpath = "quang_monorepo/foo/pkg",
    visibility = ["//visibility:public"],
    deps = ["@com_github_labstack_echo_v4//:go_default_library"],
)
