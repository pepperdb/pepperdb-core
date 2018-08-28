#ifndef _NEBULAS_NF_NVM_V8_LIB_TRACING_H_
#define _NEBULAS_NF_NVM_V8_LIB_TRACING_H_

#include <v8.h>

using namespace v8;

typedef struct {
  int source_line_offset;
  char *tracable_source;
  int strictDisallowUsage;
} TracingContext;

int InjectTracingInstructionDelegate(char **result, Isolate *isolate,
                                     const char *source, int source_line_offset,
                                     Local<Context> context, TryCatch &trycatch,
                                     void *delegateContext);

#endif // _NEBULAS_NF_NVM_V8_LIB_TRACING_H_
