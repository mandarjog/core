load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "secretfileserver.go",
        "secretserver.go",
    ],
    importpath = "istio.io/core/auth/pkg/workload",
    visibility = ["//visibility:public"],
    deps = [
        "@io_istio_auth//pkg/util:go_default_library",
    ],
)
