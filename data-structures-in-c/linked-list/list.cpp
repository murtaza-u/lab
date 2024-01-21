#include <iostream>

using namespace std;

class Node {
  public:
    int data;
    Node* next;

  Node(int data) {
    this -> data = data;
    this -> next = nullptr;
  }
};

class List {
  private:
    Node* head;
  public:
    List() {
      this -> head = nullptr;
    }

    List(int data) {
      this -> head = new Node(data);
    }

    ~List() {
      if (!head) return;

      Node* i = this -> head;
      while (i) {
        Node *temp = i;
        i = i -> next;
        free(temp);
      }
    }

    void print() {
      if (!this -> head) {
        cout << "NULL\n";
        return;
      }
      for (Node *i = this -> head; i; i = i -> next) {
        cout << i -> data << " -> ";
      }
      cout << "NULL\n";
    }

    void insert(int data) {
      Node* n = new Node(data);
      Node *i = this -> head;
      while (i -> next) i = i -> next;
      i -> next = n;
    }
};

int main() {
  List* list = new List();
  list -> insert(10);
  list -> insert(20);
  list -> insert(30);
  list -> insert(40);
  list -> insert(50);
  list -> print();
  
  return 0;
}