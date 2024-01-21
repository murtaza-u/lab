#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node *next;
} Node;

Node *head = NULL;

void insert(int data) {
  Node *n = (Node *)malloc(sizeof(Node));
  n->data = data;
  n->next = NULL;

  if (!head) {
    head = n;
    return;
  }

  Node *i = head;
  while (i->next != NULL)
    i = i->next;
  i->next = n;
}

bool search(int data) {
  if (!head) {
    return false;
  }
  for (Node *i = head; i; i = i->next) {
    if (i->data == data)
      return true;
  }
  return false;
}

int delete (int data) {
  if (!head) {
    return -1;
  }

  if (head->data == data) {
    Node *target = head;
    head = head->next;
    free(target);
    return 1;
  }

  Node *i = head;
  while (i->next && i->next->data != data) {
    i = i->next;
  }

  Node *target = i->next;
  if (!target)
    return -1;

  i->next = i->next->next;
  free(target);
  return 1;
}

void print() {
  if (!head) {
    printf("NULL\n");
    return;
  }
  for (Node *i = head; i; i = i->next) {
    printf("%d -> ", i->data);
  }
  printf("NULL\n");
}

int main() {
  insert(1);
  insert(2);
  insert(3);
  insert(4);
  print();

  int target = 4;
  int err = delete (target);
  if (err == -1) {
    fprintf(stderr, "failed to delete element %d from list\n", target);
    exit(1);
  }

  print();

  return 0;
}
