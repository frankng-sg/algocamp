import sys
import math

DEBUG_ENABLE = True
MAP_WIDTH = 16000
MAP_LENGTH = 9000 

MAP_ORIGIN_X = 0
MAP_ORIGIN_Y = 0

### Debug Functions ###
def debug_message(text):
    print(text, file=sys.stderr)
    
def debug_Pod_params(pod):
    debug_message(f"----- Turn {game_turn} -----")
    debug_message(f" GAME INPUT ")
    debug_message(f" Pod Location: ({myPod.curr_pos.x} {myPod.curr_pos.y})")
    debug_message(f" CheckPoint: ({myPod.dest_pos.x} {myPod.dest_pos.y})")
    debug_message(f" Distance To CheckPoint : ({myPod.curr_dist2dest})")
    debug_message(f" Pod Orientation: ({myPod.curr_angle2dest})")
    debug_message(f"")
    debug_message(f" FLIGHT PARAMS")
    debug_message(f" Steer Angle = {pod.steer_angle}")
    debug_message(f" Speed = {pod.speed}")
    debug_message(f" Velocity = ({pod.velocity.x},{pod.velocity.y})")

### Math Functions ###
def create_vector(vector, origin, arrow):
    vector.x = arrow.x - origin.x
    vector.y = arrow.y - origin.y

### Classes Declaration ###
class XY_Vector():
    x = 0
    y = 0

    def __init__(self):
        pass

class Polar_Vector():
    radius = 0
    theta = 0

    def __init__(self):
        pass

class CheckPoint(XY_Vector):
    RADIUS = 600

    def __init__(self):
        pass


class PodBlueprint():

    def __init__(self):

        self.MAX_ENGINE_THRUST = 100
        self.MIN_ENGINE_THRUST = 0
        self.BOOSTER_THRUST = 650
        self.SHIELD_RADIUS = 400

        self.curr_pos = XY_Vector()
        self.prev_pos = XY_Vector()
        self.goto_pos = XY_Vector()
        self.dest_pos = XY_Vector()

        self.curr_dist2dest = 0
        self.curr_angle2dest = 0

        self.prev_dist2dest = 0
        self.prev_angle2dest = 0

        self.steer_angle = 0
        self.speed = 0
        self.velocity = XY_Vector()

        self.shield = 0 # 0 = READY | N = ready in N turns
        self.boost = 1 # 1 = READY | 0 = DEPLETED
        self.thrust = 0               



    def log(self):
        self.prev_pos.x = self.curr_pos.x
        self.prev_pos.y = self.curr_pos.y
        self.prev_angle2dest = self.curr_angle2dest
        self.prev_dist2dest = self.curr_dist2dest

    def update(self):
        self.speed = abs(self.curr_dist2dest - self.prev_dist2dest)
        self.steer_angle = self.curr_angle2dest - self.prev_angle2dest
        create_vector(self.velocity, self.prev_pos, self.curr_pos)



##############################        MAIN PROGRAM       ###########################################

myPod = PodBlueprint()
oppPod = PodBlueprint()


### FIRST TURN ###
game_turn = 1
myPod.curr_pos.x, myPod.curr_pos.y, \
myPod.dest_pos.x, myPod.dest_pos.y, \
myPod.curr_dist2dest, myPod.curr_angle2dest = [int(i) for i in input().split()]

oppPod.curr_pos.x, oppPod.curr_pos.y = [int(i) for i in input().split()]  

myPod.log()
oppPod.log()

while True:

    if DEBUG_ENABLE:
        debug_Pod_params(myPod)

    ### OUTPUT ###
    myPod.goto_pos = myPod.dest_pos
    myPod.thrust = 100
    print(f"{myPod.goto_pos.x} {myPod.goto_pos.y} {myPod.thrust}")
    
    ### POST PROCESSING ###
    myPod.log()
    oppPod.log()

    ### INPUT ###
    game_turn += 1
    myPod.curr_pos.x, myPod.curr_pos.y, \
    myPod.dest_pos.x, myPod.dest_pos.y, \
    myPod.curr_dist2dest, myPod.curr_angle2dest = [int(i) for i in input().split()]
    
    oppPod.curr_pos.x, oppPod.curr_pos.y = [int(i) for i in input().split()]    

    myPod.update()
    oppPod.update()

    
