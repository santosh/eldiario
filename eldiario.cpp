#include <iostream>
#include <fstream>
using namespace std;

// get today's date, and check if entry is being written on the same date
// split by date (must have one space)
//
// check if we can use google docs REST API to change Essay of My Life document
// don't do procedural

int main(int argc, char *argv[]) {
    char date_today[25];
    char entry[100];
    string first_arg = argv[1];

    fstream diaryfile;
    if (first_arg == "-l") {
        string content;
        diaryfile.open("diary.txt");
        if (!diaryfile.fail()) {
            while (getline(diaryfile, content)) {
                cout << content << endl;
            }
        } else {
            cerr << "Error opening file." << endl;
        }
    } else {
        diaryfile.open("diary.txt", ios::app);
        
        // get todays date in form of [3rd June, 2018]
        cout << "Date e.g. [4th July, 2018]: ";
        cin.getline(date_today, 25);
        diaryfile << date_today << endl;
        cout << "How are you feeling today: ";
        cin.getline(entry, 100);
        diaryfile << entry << endl << endl;

        diaryfile.close();
        cout << "Done writing, closeing." << endl;
    }
    return 0;
}
