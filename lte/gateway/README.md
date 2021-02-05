# Bazel Build Notes

## Clang Install

Install Clang (per LLVM Website):

`sudo bash -c "$(wget -O - https://apt.llvm.org/llvm.sh)"`

`sudo update-alternatives --install /usr/bin/clang clang /usr/lib/llvm-11/bin/clang 100 --slave /usr/bin/clang++ clang++ /usr/lib/llvm-11/bin/clang++`

## Bazel Install

```shell script
sudo apt install curl gnupg
curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
sudo mv bazel.gpg /etc/apt/trusted.gpg.d/
echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
```

`sudo apt update && sudo apt install bazel`

## Setup

Edit .bazelrc:

```shell script
# Force the use of Clang for C++ builds.
build --action_env=CC=clang
build --action_env=CXX=clang++

# Define the --config=asan-libfuzzer configuration.
build:asan-libfuzzer --@rules_fuzzing//fuzzing:cc_engine=@rules_fuzzing//fuzzing/engines:libfuzzer
build:asan-libfuzzer --@rules_fuzzing//fuzzing:cc_engine_instrumentation=libfuzzer
build:asan-libfuzzer --@rules_fuzzing//fuzzing:cc_engine_sanitizer=asan
````

This was also required:

`sudo ln -s /usr/lib/llvm-11/bin/ld.lld /usr/bin/ld.lld`

Otherwise I get `invalid linker name in argument '-fuse-ld=lld'`.

And as there was an existing Docker clang version:

`sudo ln -s /usr/lib/llvm-11/bin/clang /usr/bin/clang`

## Build with Clang

`bazel run -c opt --config=asan-libfuzzer --linkopt="-fuse-ld=lld" c/oai/tasks/nas/emm/msg:attach_accept_fuzz_test_run`
