#ifndef _NEBULAS_NF_NVM_V8_ENGINE_INT_H_
#define _NEBULAS_NF_NVM_V8_ENGINE_INT_H_

#include "engine.h"

#include <v8.h>

using namespace v8;

typedef int (*ExecutionDelegate)(char **result, Isolate *isolate,
                                 const char *source, int source_line_offset,
                                 Local<Context> context, TryCatch &trycatch,
                                 void *delegateContext);

int Execute(char **result, V8Engine *e, const char *data, int source_line_offset,
            void *lcsHandler, void *gcsHandler, ExecutionDelegate delegate,
            void *delegateContext);

int ExecuteSourceDataDelegate(char **result, Isolate *isolate, const char *source,
                              int source_line_offset, Local<Context> context,
                              TryCatch &trycatch, void *delegateContext);

#endif // _NEBULAS_NF_NVM_V8_ENGINE_INT_H_
