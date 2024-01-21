#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node* next;
} Node;

Node* head = NULL;

void insertMany(int count, ...) {
  va_list ptr;
  va_start(ptr, count);

  for (int i = 0; i < count; i ++) {
	int data = va_arg(ptr, int);

	Node* n = (Node *)malloc(sizeof(Node));
	n -> data = data;
	n -> next = NULL;

	if (!head) {
	  head = n;
	  continue;
	}

	Node *temp = head;
	while (temp -> next != NULL) {
	  temp = temp -> next;
	}
	temp -> next = n;
  }

  va_end(ptr);
}

void insertOne(int data) {
  Node* n = (Node *)malloc(sizeof(Node));
  n -> data = data;
  n -> next = NULL;

  if (!head) {
	head = n;
	return;
  }

  Node *temp = head;
  while (temp -> next != NULL) {
	temp = temp -> next;
  }
  temp -> next = n;
}

void insertAfter(int target, int data) {
  if (!head) return;

  for (Node *i = head; i; i = i -> next) {
	if (i -> data == target) {
	  Node *temp = (Node*)malloc(sizeof(Node));
	  temp -> data = data;
	  temp -> next = i -> next;
	  i -> next = temp;
	}
  }
}

void print() {
  for (Node* n = head; n; n = n -> next) {
	printf("%d -> ", n -> data);
  }
  printf("NULL\n");
}

int sum() {
  int sum = 0;
  for (Node *n = head; n; n = n -> next) {
	sum += n->data;
  }
  return sum;
}

void reverse() {
  if (!head || !head -> next) return;

  Node *before = NULL, *i = head, *after;

  while (i) {
	  after = i -> next;
	  i -> next = before;
	  before = i;
	  i = after;
  }
  head = before;
}

int length() {
  int l = 0;
  for (Node *i = head; i; i = i -> next) {
	l ++;
  }
  return l;
}

Node* atIndex(int index) {
  int l = length();
  if (!head || index >= l) return NULL;

  Node *n = head;
  for (int i = 0; i < l; i ++) {
	if (i == index) {
	  return n;
	}
	n = n -> next;
  }

  return NULL;
}

void sort() {
  int n = length();
  for (int i = 0; i < n - 1; i ++) {
	for (int j = 0; j < n - 1 - i; j ++) {
	  Node *x = atIndex(j);
	  Node *y = atIndex(j + 1);

	  if (x -> data > y -> data) {
		int temp = x -> data;
		x -> data = y -> data;
		y -> data = temp;
	  }
	}
  }
}

void erase(int target) {
  Node *before = NULL;

  for (Node *i = head; i; i = i -> next) {
	if (i -> data != target) {
	  before = i;
	  continue;
	}

	if (before) {
	  before -> next = i -> next;
	}

	if (i == head) {
	  head = i -> next;
	}

	free(i);
	return;
  }
}

void printWithRecursion(Node* n) {
  if (!n) {
	printf("NULL\n");
	return;
  }
  printf("%d -> ", n -> data);
  printWithRecursion(n -> next);
}

int main() {
  printf("\n# Creating list\n");
  insertMany(5, 7, 3, 4, 9, 11);
  print();

  printf("\n# Adding 15 to it\n");
  insertOne(15);
  print();

  printf("\n# Calculating sum\n");
  printf("Sum = %d\n", sum());

  printf("\n# Reverse\n");
  reverse();
  print();

  printf("\n# Inserting 5 after 3\n");
  insertAfter(3, 5);
  print();

  printf("\n# Sorting\n");
  sort();
  print();

  printf("\n# Deleting\n");
  erase(5);
  print();

  printf("\n# Printing with recursion\n");
  printWithRecursion(head);

  return 0;
}
