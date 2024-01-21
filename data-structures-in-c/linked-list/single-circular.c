#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node *next;
} Node;

Node *head = NULL;

void print() {
  if (!head) {
	printf("NULL\n");
	return;
  }

  Node *i = head;
  printf("... -> %d -> ", i -> data);

  for (i = i -> next; i != head; i = i -> next) {
	printf("%d -> ", i -> data);
  }
  printf("...\n");
}

int length() {
  if (!head) return 0;

  int l = 1;
  for (Node *i = head -> next; i != head; i = i -> next) {
	l ++;
  }
  return l;
}

Node* atIndex(int idx) {
  int len = length();
  if (idx < 0 || idx >= len) return NULL;

  Node *n = head;
  for (int i = 1; i <= idx; i ++) {
	n = n -> next;
  }
  return n;
}

void insertOne(int data) {
  Node *n = (Node *)malloc(sizeof(Node));
  n -> data = data;

  if (!head) {
	n -> next = n;
	head = n;
	return;
  }

  Node *i = head;
  while (i -> next != head) i = i -> next;
  i -> next = n;
  n -> next = head;
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
  if (!head) return;

  Node *n = head;
  for (int i = 0; i < length(); i ++) {
	if (n -> data != target) {
	  n = n -> next;
	  continue;
	}

	Node *temp = (Node *)malloc(sizeof(Node));
	temp -> data = data;
	temp -> next = n -> next;
	n -> next = temp;
	return;
  }
}

void insertBefore(int target, int data) {
  if (!head) return;

  Node *n = head;
  for (int i = 0; i < length(); i ++) {
	if (n -> next -> data != target) {
	  n = n -> next;
	  continue;
	}

	Node *temp = (Node *)malloc(sizeof(Node));
	temp -> data = data;
	temp -> next = n -> next;
	n -> next = temp;
	return;
  }
}

void erase(int data) {
  if (!head) return;

  Node *n = head;
  for (int i = 0; i < length(); i ++) {
	if (n -> next -> data != data) {
	  n = n -> next;
	  continue;
	}

	Node *temp = n -> next;
	n -> next = n -> next -> next;

	if (temp == head) {
	  head = n -> next;
	}

	free(temp);
	return;
  }
}

int main() {
  insertMany(5, 35, -2, 54, 1, 43);
  print();
  printf("length: %d\n", length());

  erase(35);
  print();

  insertAfter(43, 100);
  print();

  insertBefore(43, 200);
  print();
}
