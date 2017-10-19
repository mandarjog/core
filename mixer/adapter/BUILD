# gazelle:ignore
package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load(":inventory.bzl", "inventory_library")

inventory_library(
    name = "go_default_library",
    packages = {
        "denier": "istio.io/core/mixer/adapter/denier",
        "list": "istio.io/core/mixer/adapter/list",
        "noop": "istio.io/core/mixer/adapter/noop",
        "prometheus": "istio.io/core/mixer/adapter/prometheus",
        "stackdriver": "istio.io/core/mixer/adapter/stackdriver",
        "statsd": "istio.io/core/mixer/adapter/statsd",
        "stdio": "istio.io/core/mixer/adapter/stdio",
        "svcctrl": "istio.io/core/mixer/adapter/svcctrl",
        "memquota": "istio.io/core/mixer/adapter/memquota",
    },
    deps = [
        "//mixer/adapter/denier:go_default_library",
        "//mixer/adapter/list:go_default_library",
        "//mixer/adapter/memquota:go_default_library",
        "//mixer/adapter/noop:go_default_library",
        "//mixer/adapter/prometheus:go_default_library",
        "//mixer/adapter/stackdriver:go_default_library",
        "//mixer/adapter/statsd:go_default_library",
        "//mixer/adapter/stdio:go_default_library",
        "//mixer/adapter/svcctrl:go_default_library",
    ],
)
