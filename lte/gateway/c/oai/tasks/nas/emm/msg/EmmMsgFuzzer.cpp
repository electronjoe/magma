#include <cstddef>
#include <cstdint>
#include <cstdio>
#include <cstdlib>


extern "C" {
#include "emm_msg.h"
}

extern "C" int LLVMFuzzerTestOneInput(const uint8_t* Data, size_t Size) {
  EMM_msg msg;
  emm_msg_decode(&msg, const_cast<uint8_t*>(Data), static_cast<uint32_t>(Size));
  return 0;  // Non-zero return values are reserved for future use.
}
