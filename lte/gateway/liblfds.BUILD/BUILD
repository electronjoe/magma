package(default_visibility = ["//visibility:public"])

cc_library(
    name = "liblfds",
    srcs = [
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_misc/lfds710_misc_internal_backoff_init.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_cleanup.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_dequeue.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_enqueue.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_init.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_query.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_cleanup.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_init.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_pop.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_push.c",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_query.c",
    ],
    hdrs = [
        "liblfds/liblfds7.1.0/liblfds710/inc/liblfds710.h",
        "liblfds/liblfds7.1.0/liblfds710/inc/liblfds710/lfds710_queue_bounded_manyproducer_manyconsumer.h",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_misc/lfds710_misc_internal.h",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_queue_bounded_manyproducer_manyconsumer/lfds710_queue_bounded_manyproducer_manyconsumer_internal.h",
        "liblfds/liblfds7.1.0/liblfds710/src/lfds710_stack/lfds710_stack_internal.h",
        "liblfds/liblfds7.1.0/liblfds710/src/liblfds710_internal.h",
    ],
    includes = [
        "liblfds/liblfds7.1.0/liblfds710/inc/",
        "liblfds/liblfds7.1.0/liblfds710/inc/liblfds710/",
        "liblfds/liblfds7.1.0/liblfds710/src/",
    ],  # Ugly but couldn't find a way around
)
