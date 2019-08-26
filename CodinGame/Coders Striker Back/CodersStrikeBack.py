import sys
import math

MAP_WIDTH = 16000
MAP_LENGTH = 9000 

MAP_ORIGIN_X = 0
MAP_ORIGIN_Y = 0

class XY_Coordiate():
    x = 0
    y = 0

    def __init__(self):
        pass

class BattleMap():



    def __init__(self):
        pass

class CheckPoint(XY_Coordiate):

    RADIUS = 600

    def __init__(self):
        pass


class Pod:

    MAX_ENGINE_THRUST = 100
    MIN_ENGINE_THRUST = 0
    BOOSTER_THRUST = 650
    SHIELD_RADIUS = 400

    curr_pos = XY_Coordiate()
    prev_pos = XY_Coordiate()
    goto_pos = XY_Coordiate()
    dest_pos = XY_Coordiate()
    
    dist_2_dest = 0
    angle_2_dest = 0
    shield = 0 # 0 = READY | N = ready in N turns
    boost = 1 # 1 = READY | 0 = DEPLETED
    thrust = 0

    def __init__(self):
        pass





##############################        MAIN PROGRAM       ###########################################

myPod = Pod()
enemy = Pod()

while True:
    ### INPUT ###
    myPod.curr_pos.x, myPod.curr_pos.y, \
    myPod.dest_pos.x, myPod.dest_pos.y, \
    myPod.dist_2_dest, myPod.angle_2_dest = [int(i) for i in input().split()]
    
    enemy.curr_pos.x, enemy.curr_pos.y = [int(i) for i in input().split()]

    
    ### OUTPUT ###
    myPod.goto_pos = myPod.dest_pos
    myPod.thrust = 100
    print(f"{myPod.goto_pos.x} {myPod.goto_pos.y} {myPod.thrust}")
