#ifndef _NEBULAS_NF_NVM_V8_LIB_MEMORY_STORAGE_H_
#define _NEBULAS_NF_NVM_V8_LIB_MEMORY_STORAGE_H_

#include <stddef.h>

void *CreateStorageHandler();
void DeleteStorageHandler(void *handler);

char *StorageGet(void *handler, const char *key, size_t *cnt);
int StoragePut(void *handler, const char *key, const char *value, size_t *cnt);
int StorageDel(void *handler, const char *key, size_t *cnt);

#endif // _NEBULAS_NF_NVM_V8_LIB_MEMORY_STORAGE_H_
