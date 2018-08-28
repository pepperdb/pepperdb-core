#include "global.h"
#include "blockchain.h"
#include "event.h"
#include "instruction_counter.h"
#include "log_callback.h"
#include "require_callback.h"
#include "storage_object.h"
#include "crypto.h"

Local<ObjectTemplate> CreateGlobalObjectTemplate(Isolate *isolate) {
  Local<ObjectTemplate> globalTpl = ObjectTemplate::New(isolate);
  globalTpl->SetInternalFieldCount(1);

  NewNativeRequireFunction(isolate, globalTpl);
  NewNativeLogFunction(isolate, globalTpl);
  NewNativeEventFunction(isolate, globalTpl);

  NewStorageType(isolate, globalTpl);

  return globalTpl;
}

void SetGlobalObjectProperties(Isolate *isolate, Local<Context> context,
                               V8Engine *e, void *lcsHandler,
                               void *gcsHandler) {
  // set e to global.
  Local<Object> global = context->Global();
  global->SetInternalField(0, External::New(isolate, e));

  NewStorageTypeInstance(isolate, context, lcsHandler, gcsHandler);
  NewInstructionCounterInstance(isolate, context,
                                &(e->stats.count_of_executed_instructions), e);
  NewBlockchainInstance(isolate, context, lcsHandler);
  NewCryptoInstance(isolate, context);
}

V8Engine *GetV8EngineInstance(Local<Context> context) {
  Local<Object> global = context->Global();
  Local<Value> val = global->GetInternalField(0);

  if (!val->IsExternal()) {
    return NULL;
  }

  return static_cast<V8Engine *>(Local<External>::Cast(val)->Value());
}
