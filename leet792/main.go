package main

import (
	"fmt"
)

// 输入: s = "abcde", words = ["a","bb","acd","ace"]
// 输出: 3
// 解释: 有三个是 s 的子序列的单词: "a", "acd", "ace"。
type Trie struct {
	End   bool
	Child map[byte]*Trie
}

func Constructor() Trie {
	res := new(Trie)
	res.Child = make(map[byte]*Trie)
	return *res
}

func (this *Trie) Insert(word string) {
	tmpTrie := this
	for i := 0; i < len(word); i++ {
		if _, ok := tmpTrie.Child[word[i]]; !ok {
			tmpNode := Constructor()
			tmpTrie.Child[word[i]] = &tmpNode
		}
		tmpTrie = tmpTrie.Child[word[i]]
	}
	tmpTrie.End = true
}

func (this *Trie) Search(word string) bool {
	tmpNode := this
	for i := 0; i < len(word); i++ {
		if val, ok := tmpNode.Child[word[i]]; ok {
			tmpNode = val
		} else {
			return false
		}
	}
	return tmpNode.End
}

func (this *Trie) StartsWith(prefix string) bool {
	tmpNode := this
	for i := 0; i < len(prefix); i++ {
		if val, ok := tmpNode.Child[prefix[i]]; ok {
			tmpNode = val
		} else {
			return false
		}
	}
	return true
}

func main() {
	res := Constructor()
	res.Insert("beer")
	res.Insert("add")
	res.Insert("jam")
	res.Insert("rental")
	fmt.Println(res.Search("apps"))
	fmt.Println(res.Search("appl"))
	fmt.Println(res.Search("ad"))
	fmt.Println(res.Search("applepie"))
	fmt.Println(res.Search("rest"))
	fmt.Println(res.Search("jan"))
	fmt.Println(res.Search("rent"))
	fmt.Println(res.Search("beer"))
	fmt.Println(res.Search("jam"))
}
