load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["token.go"],
    importpath = "istio.io/core/auth/pkg/credential",
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_cloud_go//compute/metadata:go_default_library",
    ],
)
