#include <stdarg.h>
#include <stdio.h>
#include <stdbool.h>

#define STACK_SIZE 50

int stack[STACK_SIZE];
int topIdx = -1;

void print() {
  if (topIdx == -1) {
	printf("NULL\n");
	return;
  }

  int i = topIdx;
  while (i != -1) {
	printf("%d\n", stack[i--]);
  }
}

void push(int data) {
  if (topIdx == STACK_SIZE - 1) {
	return;
  }
  stack[++topIdx] = data;
}

void pushMany(int count, ...) {
  va_list ptr;
  va_start(ptr, count);

  for (int i = 0; i < count; i ++) {
	int data = va_arg(ptr, int);
	push(data);
  }

  va_end(ptr);
}

int pop() {
  if (topIdx == -1) {
	return -1;
  }
  return stack[topIdx--];
}

int top() {
  if (topIdx == -1) {
	return -1;
  }
  return stack[topIdx];
}

bool isEmpty() {
  return topIdx == -1;
}

int length() {
  return topIdx + 1;
}

int main() {
  pushMany(5, 1, 2, 3, 4, 5);
  print();

  printf("pop(): %d\n", pop());
  printf("top(): %d\n", top());
  print();

  printf("isEmpty(): %d\n", isEmpty());

  int len = length();
  printf("length(): %d\n", len);

  for (int i = 0; i < len; i ++) {
	pop();
  }
  printf("isEmpty(): %d\n", isEmpty());
  printf("length(): %d\n", length());
}
