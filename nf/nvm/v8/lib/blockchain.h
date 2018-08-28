#ifndef _NEBULAS_NF_NVM_V8_LIB_BLOCKCHAIN_H_
#define _NEBULAS_NF_NVM_V8_LIB_BLOCKCHAIN_H_

#include <v8.h>

using namespace v8;

void NewBlockchainInstance(Isolate *isolate, Local<Context> context,
                           void *handler);

void BlockchainConstructor(const FunctionCallbackInfo<Value> &info);
void GetTransactionByHashCallback(const FunctionCallbackInfo<Value> &info);
void GetAccountStateCallback(const FunctionCallbackInfo<Value> &info);
void TransferCallback(const FunctionCallbackInfo<Value> &info);
void VerifyAddressCallback(const FunctionCallbackInfo<Value> &info);
void GetPreBlockHashCallback(const FunctionCallbackInfo<Value> &info); 
void GetPreBlockSeedCallback(const FunctionCallbackInfo<Value> &info); 


#endif //_NEBULAS_NF_NVM_V8_LIB_BLOCKCHAIN_H_
