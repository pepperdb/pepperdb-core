#ifndef _NEBULAS_NF_NVM_V8_LIB_TYPESCRIPT_H_
#define _NEBULAS_NF_NVM_V8_LIB_TYPESCRIPT_H_

#include <stddef.h>
#include <v8.h>

using namespace v8;

typedef struct {
  int source_line_offset;
  char *js_source;
} TypeScriptContext;

int TypeScriptTranspileDelegate(char **result, Isolate *isolate,
                                const char *source, int source_line_offset,
                                Local<Context> context, TryCatch &trycatch,
                                void *delegateContext);

#endif // _NEBULAS_NF_NVM_V8_LIB_TYPESCRIPT_H_
