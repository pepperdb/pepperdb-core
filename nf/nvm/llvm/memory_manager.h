#include <llvm/ExecutionEngine/SectionMemoryManager.h>
#include <string>
#include <unordered_map>

using namespace llvm;

class MemoryManager : public SectionMemoryManager {
  MemoryManager(const MemoryManager &) = delete;
  void operator=(const MemoryManager &) = delete;

public:
  MemoryManager();
  virtual ~MemoryManager();

  void addNamedFunction(const char *Name, void *Address);
  void addNamedFunction(const std::string &Name, void *Address);

  /// This method returns a RuntimeDyld::SymbolInfo for the specified function
  /// or variable. It is used to resolve symbols during module linking.
  ///
  /// By default this falls back on the legacy lookup method:
  /// 'getSymbolAddress'. The address returned by getSymbolAddress is treated as
  /// a strong, exported symbol, consistent with historical treatment by
  /// RuntimeDyld.
  ///
  /// Clients writing custom RTDyldMemoryManagers are encouraged to override
  /// this method and return a SymbolInfo with the flags set correctly. This is
  /// necessary for RuntimeDyld to correctly handle weak and non-exported
  /// symbols.
  virtual JITSymbol findSymbol(const std::string &Name);

private:
  std::unordered_map<std::string, uint64_t> namedFunctionMap;
};
