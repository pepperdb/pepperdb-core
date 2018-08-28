#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>
#include <stdint.h>

typedef struct EngineStruct {
  void *llvm_engine;
  void *llvm_builder;
  void *llvm_main_module;
  void *llvm_context;
  void *llvm_pass_manager;
  void *llvm_mem_manager;
} Engine;

Engine *CreateEngine();
int AddModuleFile(Engine *e, const char *irFile);
void DeleteEngine(Engine *e);
int RunFunction(Engine *e, const char *funcName, size_t len,
                const uint8_t *data);
void AddNamedFunction(Engine *e, const char *funcName, void *address);

void Initialize();

#ifdef __cplusplus
}
#endif
