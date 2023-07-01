#include<iostream>
using namespace std;

int main(){

	int *z;
	int *y;


	if (z==y) cout<<"Test 1";
	else {
		cout<<"Test Failed\n";	
		cout<<&z<<" First Nil Address "<<&y<<"\n";
	}
	

}
