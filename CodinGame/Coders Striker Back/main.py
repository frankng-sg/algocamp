import sys
import math


class BattleMap():

    WIDTH = 16000
    LENGTH = 9000 

    ORIGIN_X = 0
    ORIGIN_Y = 0

    def __init__(self):
        pass

class CheckPoint:

    RADIUS = 600

    x = 0
    y = 0

    def __init__(self):
        pass

class ForceField:
    
    RADIUS = 400

    def __init__(self):
        pass

class Pod:

    curr_x = 0
    curr_y = 0
    prev_x = 0
    prev_y = 0

    dest = CheckPoint()
    dist_2_dest = 0

    angle_2_dest = 0

    shield = 0 # 0 = READY | N = ready in N turns
    boost = 1 # 1 = READY | 0 = DEPLETED

    action = FlightAction()

    def __init__(self):
        pass

class FlightAction():
    
    MAX_THRUST = 100
    MIN_THRUST = 0

    goto_x = 0
    goto_y = 0
    thrust = 0

    def __init__(self):
        pass        

if __name__ = "__main__":

    myPod = Pod()
    enemy = Pod()
    
    while True:
        myPod.x, myPod.y, \
        myPod.dest.x, myPod.dest.y, \
        myPod.dist_2_dest, myPod.angle_2_dest = [int(i) for i in input().split()]
        
        enemy.curr_x, enemy.curr_y = [int(i) for i in input().split()]

        # Write an action using print
        # To debug: print("Debug messages...", file=sys.stderr)


        myPod.action.goto_x = myPod.dest.x
        myPod.action.goto_y = myPod.dest.y
        myPod.action.thrust = 50
        print(str(myPod.action.goto_x) + " " + str(myPod.action.goto_y) + myPod.action.thrust)    
    