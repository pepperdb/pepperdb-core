#ifndef _NEBULAS_NF_NVM_V8_LIB_MEMORY_MODULES_H_
#define _NEBULAS_NF_NVM_V8_LIB_MEMORY_MODULES_H_

#include <stddef.h>

char *RequireDelegateFunc(void *handler, const char *filename,
                          size_t *lineOffset);

char *AttachLibVersionDelegateFunc(void *handler, const char *libname);

void AddModule(void *handler, const char *filename, const char *source,
               int lineOffset);

char *GetModuleSource(void *handler, const char *filename);

#endif // _NEBULAS_NF_NVM_V8_LIB_MEMORY_MODULES_H_
