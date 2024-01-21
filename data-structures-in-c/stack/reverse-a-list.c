#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

/* --- Linked list --- */
typedef struct ListNode {
  int data;
  struct ListNode* next;
} ListNode;

ListNode *head = NULL;

void printList() {
  if (!head) {
	printf("NULL\n");
	return;
  }
  for (ListNode *i = head; i; i = i -> next) {
	printf("%d -> ", i -> data);
  }
  printf("NULL\n");
}

void insertIntoList(int count, ...) {
  va_list ptr;
  va_start(ptr, count);

  for (int i = 0; i < count; i ++) {
	ListNode *n = (ListNode *)malloc(sizeof(ListNode));
	n -> data = va_arg(ptr, int);
	n -> next = NULL;

	if (!head) {
	  head = n;
	  continue;
	}

	ListNode *i = head;
	while (i -> next) i = i -> next;
	i -> next = n;
  }

  va_end(ptr);
}

ListNode *popFromList() {
  if (!head) return NULL;
  ListNode *temp = head;
  head = head -> next;
  return temp;
}

void attachToList(ListNode *n) {
  n -> next = NULL;

  if (!head) {
	head = n;
	return;
  }

  ListNode *i = head;
  while (i -> next) i = i -> next;
  i -> next = n;
}

/* --- Stack --- */
typedef struct StackNode {
  ListNode* data;
  struct StackNode* next;
} StackNode;

StackNode *stack = NULL;

void printStack() {
  if (!stack) {
	printf("NULL");
	return;
  }

  for (StackNode *i = stack; i; i = i -> next) {
	printf("%d\n", i -> data -> data);
  }
}

void pushToStack(ListNode* data) {
  StackNode *n = (StackNode *)malloc(sizeof(StackNode));
  n -> data = data;
  n -> next = stack;
  stack = n;
}

ListNode* popFromStack() {
  if (!stack) return NULL;

  StackNode *temp = stack;
  stack = stack -> next;
  ListNode* data = temp -> data;
  free(temp);
  return data;
}

ListNode* top() {
  if (!stack) return NULL;
  return stack -> data;
}

bool isEmpty() {
  return stack == NULL;
}

void reverse() {
  for (ListNode* i = head; i; i = i -> next) {
	printf("pushing %d to the stack\n", i -> data);
	pushToStack(i);
  }

  while (!isEmpty()) {
	ListNode* n = popFromStack();
	printf("attaching %d to the list\n", n -> data);
	attachToList(n);
  }
}

int main() {
  insertIntoList(5, 1, 2, 3, 4, 5);

  printf("BEFORE\n");
  printList();

  reverse();

  printf("AFTER\n");
  printList();

  return 0;
}
