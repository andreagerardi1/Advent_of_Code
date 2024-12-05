#include <iostream>
#include <fstream>
#include <sstream>
#include <regex>
#include <string>
#include <vector>
using namespace std;

int main() {
    string myText;
    fstream MyFile("input.txt");
    vector<string> rows;

    while (getline(MyFile, myText)) {
        rows.push_back(myText);
    }

    int result = 0;
    regex pattern("mul\\(\\d{1,3}\\,\\d{1,3}\\)");

    for (string s : rows) {
        auto words_begin = sregex_iterator(s.begin(), s.end(), pattern);
        auto words_end = sregex_iterator();

        for (sregex_iterator i = words_begin; i != words_end; ++i) {
            smatch match = *i;
			string mult = match.str();
            int initsize = mult.size()-5;
            mult = mult.substr(4, initsize);
            int primo = stoi(mult.substr(0, mult.find(',')));
            int secondo = stoi(mult.substr(mult.find(',')+1));
            result += primo*secondo;
        }
    }
    cout << result << endl;
    return 0;
}
