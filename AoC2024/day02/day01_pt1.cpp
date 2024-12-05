#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
using namespace std;

int main() {
    string myText;
    fstream MyFile("firstinput.txt");
    int safe = 0;

    while(getline(MyFile, myText)) {
        bool increasing;
        stringstream ss(myText);
        int a, b;
        ss >> b >> a;
        increasing = b < a;
        if (b == a || abs(b-a)>3) continue;
        bool coherent = true;
        while (ss >> b) {
            if ((a == b)||(increasing && a > b) || (!increasing && a<b) || (abs(a-b) > 3) ) {
                coherent = false;
                break;
            }
            a = b;
        }
        if (coherent) {
            safe ++;
        }
    }
    cout << safe << endl;
}