#ifndef _NEBULAS_NF_NVM_V8_LIB_REQUIRE_CALLBACK_H_
#define _NEBULAS_NF_NVM_V8_LIB_REQUIRE_CALLBACK_H_

#include <v8.h>

using namespace v8;
#define LIB_WHITE   "lib/contract.js"

void NewNativeRequireFunction(Isolate *isolate,
                              Local<ObjectTemplate> globalTpl);
void RequireCallback(const v8::FunctionCallbackInfo<v8::Value> &info);

#endif // _NEBULAS_NF_NVM_V8_LIB_REQUIRE_CALLBACK_H_
