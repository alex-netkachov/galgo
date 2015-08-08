#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>
#include <assert.h>


void insertionSort(int N, int arr[]) {
    int shifts = 0;
    int i,j;
    int value;
    for(i=1;i<N;i++)
    {
        value=arr[i];
        j=i-1;
        while(j>-1 && value<arr[j])
        {
            shifts++;
            arr[j+1]=arr[j];
            j=j-1;
        }
        arr[j+1]=value;
    }
    printf("%d",shifts);
}
int main(void) {

    int N;
    scanf("%d", &N);
    int arr[N], i;
    for(i = 0; i < N; i++) {
        scanf("%d", &arr[i]);
    }

    insertionSort(N, arr);

    return 0;
}