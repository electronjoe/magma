load("@rules_fuzzing//fuzzing:cc_defs.bzl", "cc_fuzz_test")

cc_fuzz_test(
    name = "attach_accept_fuzz_test",
    srcs = ["AttachAcceptFuzzer.cpp"],
    deps = [
        ":attach_accept",
    ],
)

cc_library(
    name = "attach_accept",
    srcs = [
        "AttachAccept.c",
    ],
    hdrs = [
        "AttachAccept.h",
    ],
    deps = [
        "//c/oai/common",
        "//c/oai/lib/bstr",
        "//c/oai/tasks/nas/ies",
    ],
)

cc_fuzz_test(
    name = "emm_msg_fuzz_test",
    srcs = ["EmmMsgFuzzer.cpp"],
    deps = [
        ":emm_msg",
    ],
)

cc_library(
    name = "emm_msg",
    srcs = [
        "emm_msg.c",
    ],
    hdrs = [
        "emm_msg.h",
        "emm_msgDef.h",
    ],
    deps = [
        "//c/oai/common",
        "//c/oai/tasks/nas/ies",
    ],
)
