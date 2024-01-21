#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
  int data;
  struct Node *left;
  struct Node *right;
} Node;

Node *root = NULL;
int count = 0;

void insert(int data, Node *root) {
  if (!root) {
    root = (Node *)malloc(sizeof(Node));
    root->data = data;
    root->left = NULL;
    root->right = NULL;
    count++;
    return;
  }

  if (root->data < data) {
    if (root->left) {
      insert(data, root->left);
      return;
    }

    root->left = (Node *)malloc(sizeof(Node));
    root->left->data = data;
    root->left->left = NULL;
    root->left->right = NULL;
    count++;
    return;
  }

  if (root->right) {
    insert(data, root->right);
    return;
  }

  root->right = (Node *)malloc(sizeof(Node));
  root->right->data = data;
  root->right->left = NULL;
  root->right->right = NULL;
  count++;
}

int main() {
  insert(10, root);
  insert(20, root);
  insert(30, root);
  insert(40, root);

  printf("total elements: %d\n", count);

  return 0;
}
