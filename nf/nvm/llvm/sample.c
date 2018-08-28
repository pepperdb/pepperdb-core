// This is a sample smart contract written in C.

#include <stdio.h>

int roll_dice();

void func_a() { printf("called to func_a.\n"); }

void func_b() {
  int v = roll_dice();
  printf("called to func_b, dice is %d.\n", v);
}

int main(int argc, char *argv[]) {
  printf("called to main.\n");
  printf("argc = %d\n", argc);
  func_a();
  func_b();
  return 234;
}
