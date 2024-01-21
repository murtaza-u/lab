#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Node {
  char data;
  struct Node* next;
} Node;

Node *head = NULL;

void print() {
  if (!head) {
	printf("NULL\n");
	return;
  }

  for (Node *i = head; i; i = i -> next) {
	printf("%c\n", i -> data);
  }
}

int length() {
  int l = 0;
  for (Node *i = head; i; i = i -> next) {
	l ++;
  }
  return l;
}

void push(char data) {
  Node *n = (Node *)malloc(sizeof(Node));
  n -> data = data;
  n -> next = head;
  head = n;
}

char pop() {
  if (!head) return -1;

  Node *temp = head;
  head = head -> next;
  char data = temp -> data;
  free(temp);
  return data;
}

char top() {
  if (!head) return -1;
  return head -> data;
}

void reverse(char *str) {
  int len = strlen(str);
  for (int i = 0;  i < len; i ++) {
	push(str[i]);
  }

  for (int i = 0; i < len; i ++) {
	str[i] = pop();
  }
}

int main() {
  char str[100] = "mana";

  printf("Before: %s\n", str);
  reverse(str);
  printf("After: %s\n", str);

  return 0;
}
