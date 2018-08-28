#ifndef _NEBULAS_NF_NVM_V8_LIB_GLOBAL_H_
#define _NEBULAS_NF_NVM_V8_LIB_GLOBAL_H_

#include "../engine.h"
#include <v8.h>

using namespace v8;

Local<ObjectTemplate> CreateGlobalObjectTemplate(Isolate *isolate);

void SetGlobalObjectProperties(Isolate *isolate, Local<Context> context,
                               V8Engine *e, void *lcsHandler, void *gcsHandler);

V8Engine *GetV8EngineInstance(Local<Context> context);

#endif // _NEBULAS_NF_NVM_V8_LIB_GLOBAL_H_
