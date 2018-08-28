#include "engine.h"
#include <stdio.h>
#include <stdlib.h>

void LogError(const char *msg) { printf("Error: %s\n", msg); }

int roll_dice() { return 233; }

void help(const char *name) {
  printf("%s [IR file] [Func Name] [Args...]\n", name);
  exit(1);
}

int main(int argc, const char *argv[]) {
  if (argc < 3) {
    help(argv[0]);
  }

  const char *filePath = argv[1];
  const char *funcName = argv[2];

  Initialize();
  printf("initialized.\n");

  Engine *e = CreateEngine();
  printf("engine created.\n");

  if (AddModuleFile(e, filePath) > 0) {
    printf("failed to add module file: %s\n", filePath);
    return 1;
  }

  printf("add module file: %s.\n", filePath);

  AddNamedFunction(e, "roll_dice", (void *)roll_dice);
  AddNamedFunction(e, "roll_dice2", (void *)roll_dice);
  printf("add named functions\n");

  int ret = RunFunction(e, funcName, 111, NULL);
  printf("runFunction return %d\n", ret);

  DeleteEngine(e);
  printf("engine deleted.\n");

  return ret;
}
