#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
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



    int similarity = 0;

    for (int i = 0; i < firstList.size(); i++) {
        int num = firstList[i];
        int occurrencies = 0;
        cout << occurrencies << endl;
        for (int j = 0; j < secondList.size(); j++) {
            if (secondList[j] == num) {
                occurrencies += 1;
            }
        }
        similarity += (num * occurrencies);
    }

    cout << similarity << endl;
    MyFile.close();
}