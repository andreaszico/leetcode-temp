class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function getNumber(temp: ListNode): number {
  let result: string = "";

  while (temp!.val != null) {
    result += temp!.val;
    if (temp!.next == null) break;
    temp = temp?.next!;
  }

  return +result.split("").reverse().join("");
}

function changeToListNode(number: number): ListNode {
  let temp: ListNode;
  let result: ListNode = new ListNode();
  let arr = number.toString().split("").reverse();

  let i = 0;
  temp = new ListNode(+arr[i]);
  while (arr.length >= i) {
    temp.next = new ListNode(+arr[i]);

    result = temp;
    temp = result;
    i++;
  }

  return result;
}

function addTwoNumbers(
  l1: ListNode | null,
  l2: ListNode | null
): ListNode | null {
  let temp = l1;

  const l1Num: number = getNumber(l1!);
  const l2Num: number = getNumber(l2!);

  return changeToListNode(l1Num + l2Num);
}

let l1 = new ListNode(2);
l1.next = new ListNode(4);
l1.next.next = new ListNode(3);

let l2 = new ListNode(5);
l2.next = new ListNode(4);
l2.next.next = new ListNode(6);

console.log(addTwoNumbers(l1, l2));
