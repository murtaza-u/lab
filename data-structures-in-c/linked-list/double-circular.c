#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node *next;
  struct Node *prev;
} Node;

Node *head = NULL;

void print() {
  if (!head) {
	printf("NULL\n");
  }

  Node *i = head;
  printf("... <-> %d <-> ", i -> data);

  for (Node *i = head -> next; i != head; i = i -> next) {
	printf("%d <-> ", i -> data);
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

void insertOne(int data) {
  Node *n = (Node *)malloc(sizeof(Node));
  n -> data = data;

  if (!head) {
	head = n;
	n -> next = n;
	n -> prev = n;
	return;
  }

  Node *tail = head -> prev;
  n -> next = head;
  n -> prev = tail;
  tail -> next = n;
  head -> prev = n;
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
	temp -> prev = n;
	temp -> next = n -> next;
	n -> next -> prev = temp;
	n -> next = temp;
	return;
  }
}

void insertBefore(int target, int data) {
  if (!head) return;

  Node *n = head;
  for (int i = 0; i < length(); i ++) {
	if (n -> data != target) {
	  n = n -> next;
	  continue;
	}

	Node *temp = (Node *)malloc(sizeof(Node));
	temp -> data = data;
	temp -> next = n;
	temp -> prev = n -> prev;
	n -> prev -> next = temp;
	n -> prev = temp;
	return;
  }
}

void erase(int data) {
  Node *n = head;
  for (int i = 0; i < length(); i ++) {
	if (n -> data != data) {
	  n = n -> next;
	  continue;
	}

	n -> prev -> next = n -> next;
	n -> next -> prev = n -> prev;
	if (n == head) {
	  head = n -> next;
	}
	free(n);
	return;
  }
}

int main() {
  insertMany(5, 1, 2, 3, 4, 5);
  print();

  printf("length: %d\n", length());

  erase(5);
  print();

  insertAfter(4, 100);
  print();

  insertBefore(100, 200);
  print();

  return 0;
}
