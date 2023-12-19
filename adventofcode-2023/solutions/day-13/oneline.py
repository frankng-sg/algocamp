import itertools as it

input_file = 'input.txt'
q = { 13: open(input_file).read().strip() }
maps = [[list(row) for row in m] for m in [m.split('\n') for m in q[13].strip().split('\n\n')]]
print('Day 13 Part 1:',((f:=lambda m:(a:=0,[(s:=min(c,len(m[0])-c),n:=1,[(r[c-s:c]!=r[c:c+s][::-1]) and (n:=0) for r in it.takewhile(lambda _:n==1,m)]) and (n and (a:=c)) for c in range(1,len(m[0]))]) and a),([[list(r) for r in m] for m in [m.split('\n') for m in q[13].strip().split('\n\n')]]),x:=0,y:=0,[(v:=f(m),h:=f(list(zip(*m))),(v and (x:=x+v)),(h and (y:=y+h*100))) for m in [[list(r) for r in m] for m in [m.split('\n') for m in q[13].strip().split('\n\n')]]]) and x+y)
print('Day 13 Part 2:',((f:=lambda m:(a:=0,[(s:=min(c,len(m[0])-c),n:=0,[[(l!=r and (n:=n+1)) for l,r in zip(r[c-s:c],r[c:c+s][::-1])] for r in m if n<=1],(n==1 and (a:=c))) for c in range(1,len(m[0]))]) and a),([[list(r) for r in m] for m in [m.split('\n') for m in q[13].strip().split('\n\n')]]),x:=0,y:=0,[(v:=f(m),h:=f(list(zip(*m))),(v and (x:=x+v)),(h and (y:=y+h*100))) for m in [[list(r) for r in m] for m in [m.split('\n') for m in q[13].strip().split('\n\n')]]]) and x+y)