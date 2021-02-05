#include <cstddef>
#include <cstdint>
#include <cstdio>
#include <cstdlib>


extern "C" {
#include "AttachAccept.h"
#include "dynamic_memory_check.h"
}

extern "C" int LLVMFuzzerTestOneInput(const uint8_t* Data, size_t Size) {
  attach_accept_msg result;
  result.esmmessagecontainer = nullptr;
  decode_attach_accept(&result, const_cast<uint8_t*>(Data), static_cast<uint32_t>(Size));
  bdestroy_wrapper(&result.esmmessagecontainer);
  return 0;  // Non-zero return values are reserved for future use.
}
