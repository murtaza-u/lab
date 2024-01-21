#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node* next;
  struct Node* prev;
} Node;

Node *head = NULL;

void print() {
  if (!head) {
	printf("NULL\n");
	return;
  }

  printf("NULL <-> ");
  for (Node *i = head; i; i = i -> next) {
	printf("%d <-> ", i -> data);
  }
  printf("NULL\n");
}

int length() {
  int l = 0;
  for (Node *i = head; i; i = i -> next) l ++;
  return l;
}

Node* atIndex(int idx) {
  int len = length();
  if (idx < 0 || idx >= len) {
	return NULL;
  }

  Node *n = head;
  for (int i = 1; i <= idx; i ++) n = n -> next;
  return n;
}

void insertOne(int data) {
  Node* n = (Node *)malloc(sizeof(Node));
  n -> data = data;
  n -> next = NULL;

  if (!head) {
	n -> prev = NULL;
	head = n;
	return;
  }

  Node *temp = head;
  // get to the last node
  while (temp -> next) temp = temp -> next;
  temp -> next = n;
  n -> prev = temp;
}

void insertMany(int count, ...) {
  va_list ptr;
  va_start(ptr, count);

  for (int i = 0; i < count; i ++) {
	int data = va_arg(ptr, int);
	insertOne(data);
  }

  va_end(ptr);
}

void insertAfter(int target, int data) {
  for (Node *i = head; i; i = i -> next) {
	if (i -> data != target) {
	  continue;
	}

	Node* n = (Node*)malloc(sizeof(Node));
	n -> data = data;
	n -> next = i -> next;
	n -> prev = i;
	if (i -> next) i -> next -> prev = n;
	i -> next = n;
	return;
  }
}

void insertBefore(int target, int data) {
  for (Node *i = head; i; i = i -> next) {
	if (i -> data != target) continue;

	Node* n = (Node*)malloc(sizeof(Node));
	n -> data = data;
	n -> next = i;
	n -> prev = i -> prev;

	if (i -> prev) i -> prev -> next = n;
	i -> prev = n;

	if (!n -> prev) head = n;
	return;
  }
}

void insertAt(int idx, int data) {
  int len = length();
  if (idx < 0 || idx >= len) {
	return;
  }

  Node *target = head;
  for (int i = 1; i <= idx; i ++) {
	target = target -> next;
  }

  Node* n = (Node*)malloc(sizeof(Node));
  n -> data = data;
  n -> next = target;
  n -> prev = target -> prev;

  if (target -> prev) target -> prev -> next = n;
  target -> prev = n;

  if (!n -> prev) head = n;
}

void erase(int data) {
  if (!head) return;

  for (Node *i = head; i; i = i -> next) {
	if (i -> data != data) continue;

	if (i -> prev) {
	  i -> prev -> next = i -> next;
	}
	if (i -> next) {
	  i -> next -> prev = i -> prev;
	}

	if (!i -> prev) {
	  head = i -> next;
	}

	free(i);
	return;
  }
}

int main() {
  insertMany(5, 35, 1, 40, 4, -2);
  print();
  printf("Node[2]: %d\n", atIndex(2) -> data);
  printf("Node[4]: %d\n", atIndex(4) -> data);
  printf("Node[0]: %d\n", atIndex(0) -> data);
  erase(40);
  print();

  /* insertAfter(1, 100); */
  /* insertBefore(-2, 200); */
  /* insertAt(4, 300); */

  print();
}
