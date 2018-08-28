#ifndef _NEBULAS_NF_NVM_V8_ENGINE_H_
#define _NEBULAS_NF_NVM_V8_ENGINE_H_
// #include <v8.h>

#if BUILDING_DLL
#define EXPORT __attribute__((__visibility__("default")))
#else
#define EXPORT
#endif

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

#include <stddef.h>
#include <stdint.h>
#include <string.h>
#include <stdbool.h>
#include "lib/nvm_error.h"

enum LogLevel {
  DEBUG = 1,
  WARN = 2,
  INFO = 3,
  ERROR = 4,
};

enum OptType {
  INSTRUCTION     = 1,
  INSTRUCTIONTS  = 2,
  RUNSCRIPT       = 3,
};

// log
typedef void (*LogFunc)(int level, const char *msg);
EXPORT const char *GetLogLevelText(int level);
EXPORT void InitializeLogger(LogFunc f);

// event.
typedef void (*EventTriggerFunc)(void *handler, const char *topic,
                                 const char *data, size_t *counterVal);
EXPORT void InitializeEvent(EventTriggerFunc trigger);

// storage
typedef char *(*StorageGetFunc)(void *handler, const char *key,
                                size_t *counterVal);
typedef int (*StoragePutFunc)(void *handler, const char *key, const char *value,
                              size_t *counterVal);
typedef int (*StorageDelFunc)(void *handler, const char *key,
                              size_t *counterVal);
EXPORT void InitializeStorage(StorageGetFunc get, StoragePutFunc put,
                              StorageDelFunc del);

// blockchain
typedef char *(*GetTxByHashFunc)(void *handler, const char *hash,
                                 size_t *counterVal);
typedef int (*GetAccountStateFunc)(void *handler, const char *address,
                                     size_t *counterVal, char **result, char **info);
typedef int (*TransferFunc)(void *handler, const char *to, const char *value,
                            size_t *counterVal);
typedef int (*VerifyAddressFunc)(void *handler, const char *address,
                                 size_t *counterVal);
typedef int (*GetPreBlockHashFunc)(void *handler, unsigned long long offset, size_t *counterVal, char **result, char **info);

typedef int (*GetPreBlockSeedFunc)(void *handler, unsigned long long offset, size_t *counterVal, char **result, char **info);



EXPORT void InitializeBlockchain(GetTxByHashFunc getTx,
                                 GetAccountStateFunc getAccount,
                                 TransferFunc transfer,
                                 VerifyAddressFunc verifyAddress,
                                 GetPreBlockHashFunc getPreBlockHash,
                                 GetPreBlockSeedFunc getPreBlockSeed);

// crypto
typedef char *(*Sha256Func)(const char *data, size_t *counterVal);
typedef char *(*Sha3256Func)(const char *data, size_t *counterVal);
typedef char *(*Ripemd160Func)(const char *data, size_t *counterVal);
typedef char *(*RecoverAddressFunc)(int alg, const char *data, const char *sign,
                                 size_t *counterVal);
typedef char *(*Md5Func)(const char *data, size_t *counterVal);
typedef char *(*Base64Func)(const char *data, size_t *counterVal);

EXPORT void InitializeCrypto(Sha256Func sha256,
                                 Sha3256Func sha3256,
                                 Ripemd160Func ripemd160,
                                 RecoverAddressFunc recoverAddress,
                                 Md5Func md5,
                                 Base64Func base64);

// version
EXPORT char *GetV8Version();

// require callback.
typedef char *(*RequireDelegate)(void *handler, const char *filename,
                                 size_t *lineOffset);
typedef char *(*AttachLibVersionDelegate)(void *handler, const char *libname);

EXPORT void InitializeRequireDelegate(RequireDelegate delegate, AttachLibVersionDelegate libDelegate);

EXPORT void InitializeExecutionEnvDelegate(AttachLibVersionDelegate libDelegate);

typedef struct V8EngineStats {
  size_t count_of_executed_instructions;
  size_t total_memory_size;
  size_t total_heap_size;
  size_t total_heap_size_executable;
  size_t total_physical_size;
  size_t total_available_size;
  size_t used_heap_size;
  size_t heap_size_limit;
  size_t malloced_memory;
  size_t peak_malloced_memory;
  size_t total_array_buffer_size;
  size_t peak_array_buffer_size;
} V8EngineStats;

typedef struct V8Engine {

  void *isolate;
  void *allocator;
  size_t limits_of_executed_instructions;
  size_t limits_of_total_memory_size;
  bool is_requested_terminate_execution;
  bool is_unexpected_error_happen;
  int testing;
  int timeout;
  
  V8EngineStats stats;
 
} V8Engine;
typedef struct v8ThreadContextInput {
  uintptr_t lcs;  
  uintptr_t gcs;
  enum OptType opt;  
  int line_offset;
  int allow_usage;
  const char *source;
} v8ThreadContextInput;
typedef struct v8ThreadContextOutput {
  int ret;  //output
  int line_offset;
  char *result; //output
} v8ThreadContextOutput;
typedef struct v8ThreadContext_ {
  V8Engine *e; 
  v8ThreadContextInput input;
  v8ThreadContextOutput output;
  bool is_finished;
} v8ThreadContext;

EXPORT void Initialize();
EXPORT void Dispose();

EXPORT V8Engine *CreateEngine();

EXPORT int RunScriptSource(char **result, V8Engine *e, const char *source,
                           int source_line_offset, uintptr_t lcsHandler,
                           uintptr_t gcsHandler);

EXPORT char *InjectTracingInstructions(V8Engine *e, const char *source,
                                       int *source_line_offset,
                                       int strictDisallowUsage);

EXPORT char *TranspileTypeScriptModule(V8Engine *e, const char *source,
                                       int *source_line_offset);

EXPORT int IsEngineLimitsExceeded(V8Engine *e);

EXPORT void ReadMemoryStatistics(V8Engine *e);

EXPORT void TerminateExecution(V8Engine *e);

EXPORT void DeleteEngine(V8Engine *e);

EXPORT void ExecuteLoop(const char *file);

EXPORT char *InjectTracingInstructionsThread(V8Engine *e, const char *source,
                                int *source_line_offset,
                                int allow_usage);
EXPORT char *TranspileTypeScriptModuleThread(V8Engine *e, const char *source,
                                int *source_line_offset);
EXPORT int RunScriptSourceThread(char **result, V8Engine *e, const char *source,
                    int source_line_offset, uintptr_t lcs_handler,
                    uintptr_t gcs_handler);

bool CreateScriptThread(v8ThreadContext *pc);
void SetRunScriptArgs(v8ThreadContext *pc, V8Engine *e, int opt, const char *source, int line_offset, int allow_usage);
#ifdef __cplusplus
}
#endif // __cplusplus

#endif // _NEBULAS_NF_NVM_V8_ENGINE_H_
