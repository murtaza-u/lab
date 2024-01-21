#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct Node {
  int data;
  struct Node* next;
} Node;

Node *head = NULL;

void print() {
  if (!head) {
	printf("NULL\n");
	return;
  }

  for (Node *i = head; i; i = i -> next) {
	printf("%d\n", i -> data);
  }
}

int length() {
  int l = 0;
  for (Node *i = head; i; i = i -> next) {
	l ++;
  }
  return l;
}

void push(int data) {
  Node *n = (Node *)malloc(sizeof(Node));
  n -> data = data;
  n -> next = head;
  head = n;
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
  if (!head) return -1;

  Node *temp = head;
  head = head -> next;
  int data = temp -> data;
  free(temp);
  return data;
}

int top() {
  if (!head) return -1;
  return head -> data;
}

bool isEmpty() {
  return head == NULL;
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
