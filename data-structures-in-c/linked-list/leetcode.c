#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

void reverse(struct ListNode* head) {
    if (!head || !head -> next) return;

    struct ListNode *before = NULL, *i = head, *after;

    while (i) {
        after = i -> next;
        i -> next = before;
        before = i;
        i = after;
    }
    head = before;
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    reverse(l1);
    reverse(l2);

    int x = 0, y = 0;

    // l1
    while (l1) {
        x = x * 10 + l1 -> val;
        l1 = l1 -> next;
    }

    // l2
    while (l2) {
        y = y * 10 + l2 -> val;
        l2 = l2 -> next;
    }

    int sum = x + y;

    struct ListNode* head = NULL;

    while (sum != 0) {
        int rem = sum % 10;
        sum /= 10;

        struct ListNode* temp = (struct ListNode*)malloc(sizeof(struct ListNode));
        temp -> next = NULL;
        temp -> val = rem;

        if (!head) {
            head = temp;
            continue;
        }

        struct ListNode *i = head;
        while (i -> next) i = i -> next;
        i -> next = temp;
    }

    return head;
}

int main() {

    return 0;
}