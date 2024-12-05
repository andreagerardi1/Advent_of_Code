#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
using namespace std;


bool checkSafe(vector<int> row) {
    bool coherent =true;
    bool increasing = row[0] < row[1];
    if (row[0] == row[1] || abs(row[0]-row[1]) >3) return false;
    for (int i = 1; i < row.size()-1; i++) {
        int a = row[i];
        int b = row[i+1];
        if ((a == b)||(increasing && a > b) || (!increasing && a<b) || (abs(a-b) > 3) ) {
                coherent = false;
                break;
            }
        
    }
    return coherent;
}

int main() {
    string myText;
    fstream MyFile("firstinput.txt");
    int safe = 0;

    while (getline(MyFile, myText)) {
        vector<int> row;
        int element;
        stringstream ss(myText);
        while (ss >> element) {
            row.push_back(element);
        }
        // row <- tutti gli elementi di una riga

        bool b = checkSafe(row);
        if (b) {
            safe++;
            continue;
        }
        for (int i = 0; i < row.size(); i++) {
            vector<int> copy;
            for (int j = 0; j < row.size(); j++) {
                if (j != i) {
                    copy.push_back(row[j]);
                }
            }
            b = checkSafe(copy);
            if (b) {
                safe++;
                break;
            }    
        }
    }
    cout << safe << endl;
}