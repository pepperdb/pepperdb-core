#ifndef _NEBULAS_NF_NVM_V8_LIB_STORAGE_OBJECT_H_
#define _NEBULAS_NF_NVM_V8_LIB_STORAGE_OBJECT_H_

#include <v8.h>

using namespace v8;

void NewStorageType(Isolate *isolate, Local<ObjectTemplate> globalTpl);
void NewStorageTypeInstance(Isolate *isolate, Local<Context> context,
                            void *lcsHandler, void *gcsHandler);

void StorageConstructor(const FunctionCallbackInfo<Value> &info);
void StorageGetCallback(const FunctionCallbackInfo<Value> &info);
void StoragePutCallback(const FunctionCallbackInfo<Value> &info);
void StorageDelCallback(const FunctionCallbackInfo<Value> &info);

#endif // _NEBULAS_NF_NVM_V8_LIB_STORAGE_OBJECT_H_
