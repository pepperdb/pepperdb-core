#include "memory_manager.h"
#include <llvm/ExecutionEngine/SectionMemoryManager.h>

MemoryManager::MemoryManager() {}

MemoryManager::~MemoryManager() {}

JITSymbol MemoryManager::findSymbol(const std::string &Name) {
  const char *NameStr = Name.c_str();

// DynamicLibrary::SearchForAddresOfSymbol expects an unmangled 'C' symbol
// name so ff we're on Darwin, strip the leading '_' off.
#ifdef __APPLE__
  if (NameStr[0] == '_') {
    ++NameStr;
  }
#endif

  printf("Name is %s, NameStr is %s\n", Name.c_str(), NameStr);

  uint64_t addr = 0;

  auto it = this->namedFunctionMap.find(std::string(NameStr));
  if (it == this->namedFunctionMap.end()) {
    addr = getSymbolAddress(Name);
  } else {
    addr = it->second;
  }

  return JITSymbol(addr, JITSymbolFlags::Exported);
}

void MemoryManager::addNamedFunction(const char *Name, void *Address) {
  this->addNamedFunction(std::string(Name), Address);
}

void MemoryManager::addNamedFunction(const std::string &Name, void *Address) {
  this->namedFunctionMap[Name] = (uint64_t)Address;
}
