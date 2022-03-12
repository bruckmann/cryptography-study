# How does this cipher work

This cipher use a tecnical called `substitution technique`, basically we have two alphabets, one normal and the other with different order, the order of our second alphabet can be with the letters 3 positions ahead for example.

## Example

This will be our normal alphabet:
```code
  A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
```
And this well be the `secret-key` alphabet:
```code
 D E D E F G H I J K L M N O P Q R S T U V W A B C
```
So when we go to encrypt a sentence for example: `test caesar cipher`, the output will be `whvw wkh fdhvdu flskhu`

## For more infos:

- [Wikipedia](https://en.wikipedia.org/wiki/Caesar_cipher)
- [GeekForGeeks](https://www.geeksforgeeks.org/caesar-cipher-in-cryptography/)