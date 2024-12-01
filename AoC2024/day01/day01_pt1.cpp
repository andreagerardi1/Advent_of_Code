#include <iostream>
#include <fstream>
#include <vector>
#include <sstream>
#include <algorithm>
using namespace std;

int main() {
    string myText;
    vector<int> firstList;
    vector<int> secondList;

    fstream MyFile("first_input.txt");
    
    while (getline(MyFile,myText)) {
        stringstream ss(myText);
        int first, second;
        ss >> first >> second;
        firstList.push_back(first);
        secondList.push_back(second);
    }

    sort(firstList.begin(), firstList.end());
    sort(secondList.begin(), secondList.end());

    int distanza;

    for (int i = 0; i < firstList.size(); i++) {
        int a, b;
        if (firstList[i] >= secondList[i]) {
            a = firstList[i];
            b = secondList[i];
        } else {
            a = secondList[i];
            b = firstList[i];
        }
        distanza += a-b;
    }

    cout << distanza << endl;

}