import time

def fibo(n):
    if n in (0, 1):
        return 1
    else:
        return fibo(n-1) + fibo(n-2)

if __name__ == '__main__':
    start = time.time()
    print(fibo(39))
    end = time.time()
    print(end - start)