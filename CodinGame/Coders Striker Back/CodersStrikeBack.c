#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <math.h>


/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
 
/*--------------------------GAME FORMULAE------------------------------------------------------------------------------
<1> The game is played on a map 16000 units wide and 9000 units high. 
	The coordinate X=0, Y=0 is the top left pixel.

<2> 0 ≤ thrust ≤ 100
	Response time first turn ≤ 1000ms
	Response time per turn ≤ 150ms

<3> To pass a checkpoint, the center of a pod must be inside the radius of the checkpoint.

<4> MAX_ENGINE_THRUST 100

<5> BOOSTER_THRUST 650
---------------------------------------------------------------------------------------------------------------------*/
 

/*--------------------------WINNING GAME PARAMETERS--------------------------------------------------------------------*/

// maximum possible number of checkpoints
#define MAX_CHECKPOINTS 5
// maxium possible number of laps
#define MAX_LAPS 3
// checkpoint radius 
// to pass a checkpoint, the center of a pod must be inside the radius of the checkpoint.
#define CHECKPOINT_RADIUS 600
#define BUFFER_RADIUS 200
#define INBOUND_RADIUS 2000
#define MIN_APPROACH_SPEED 300
#define MAX_APPROACH_SPEED 400
// minimum distance to use BOOST
#define LONG_DISTANCE 9000
#define BOOST "BOOST"
#define FULL_THROTTLE "100"
#define HALF_THROTTLE "80"
#define BRAKE "40"
#define MAX_ORIENTATION 45
#define PI 3.14
#define APPROACH_ANGLE 45
/*---------------------------------------------------------------------------------------------------------------------*/

/*--------------------------DATA STRUCTURES----------------------------------------------------------------------------*/
struct vectorType
{
	int  dx, dy;	
};
typedef struct vectorType vector;

struct coordinateType
{
	int  x, y;
};
typedef struct coordinateType planeXY;

struct podObject
{
	//Output
	planeXY goToXY; 
	char thrustPower[6];
	//Status
	char booster; // 1 = Availabe    0 = Empty
	int speed;
	int orientation; // angle between pod orientation and the line connect pod and checkpoint
	int direction; // angle between pod velocity and the line connect pod and checkpoint
	int distance; // the distance between pod center to center of checkpoint
	planeXY prevLoc, currLoc;
	vector velocity;	
};
typedef struct podObject pod;
/*---------------------------------------------------------------------------------------------------------------------*/

/*--------------------------MAP VARIABLES------------------------------------------------------------------------------*/

// store array of checkpoint coordinates (x, y) on the map
planeXY listCheckpoint[MAX_CHECKPOINTS];
// store number of checkpoints on the map
int nCheckpoints;
// store the index of current checkpoint in array
int cpIndex;
// store the index of current racing lap
int lapIndex;
// 
char lapFlag;

/*---------------------------------------------------------------------------------------------------------------------*/


/*--------------------------MY POD VARIABLES---------------------------------------------------------------------------*/
pod myPod;
/*---------------------------------------------------------------------------------------------------------------------*/


/*--------------------------OPPONENT POD VARIABLES---------------------------------------------------------------------*/
pod opPod;
/*---------------------------------------------------------------------------------------------------------------------*/

/*-------------------------MAIN FUNCTION PROTOTYPE---------------------------------------------------------------------*/
void initialize(void);
void computeMapVariables(int checkpointX, int checkpointY);		// Input = current checkpoint coordinate   
																// Output = lapIndex, cpIndex, nCheckpoints    
void computeMypodVariables(int xo, int yo, int dist, int angle);// Input = current location (xo, yo),  distance to checkpoint (dist), 
																//			angle between the orientation and displacement to checkpoint (angle) in degree 
																// Output = mypodVelocity, mypodOrientation, mypodDistance
void computeOpponentVariables(int xo, int yo); 	// Input = current opponent coordinate
												// Output = oppCurrLoc, oppCurrVelocity, oppCurrAcceleration, oppCurrDistance, oppAvailBoost
void followOpponent(void); // To test 
void followCheckpoint(void); // To test angle and velocity
/*---------------------------------------------------------------------------------------------------------------------*/

/*-------------------------SUPPORT FUNCTION PROTOTYPE------------------------------------------------------------------*/
void addCheckpoint(int xo, int yo);
char recognizedCheckpoint(int xo, int yo);
double toRad(double angleInDegree);
int toDegree(double angleInRad);
int vectorLength(vector u);
double dotProduct(vector u, vector v);
vector subtractVectors(vector u, vector v);
vector addVectors(vector u, vector v);
void toString(char str[6], int n);
void myPodGoTo(planeXY destination, char * thrust);
// create a vector from two points
vector createVector(planeXY pU, planeXY pV);
int calcAngleBetweenVectors(vector u, vector v); //return angle in degree
char possibleForCornerTurn();
char requireBoostingSpeed();
void optimizeApproachingSpeed();
/*---------------------------------------------------------------------------------------------------------------------*/

int main()
{
    //Initialize the game parameters
    initialize();
    // game loop
    while (1) 
	{
   		
    	//Reading game inputs
        int x;
        int y;
        int nextCheckpointX; // x position of the next check point
        int nextCheckpointY; // y position of the next check point
        int nextCheckpointDist; // distance to the next checkpoint
        int nextCheckpointAngle; // angle in degree between your pod orientation and the direction of the next checkpoint
        scanf("%d%d%d%d%d%d", &x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle);
        int opponentX;
        int opponentY;
        scanf("%d%d", &opponentX, &opponentY);

		/*----computation of all parameters derived from game inputs----------------------------------------------------*/        
		
        computeMapVariables(nextCheckpointX, nextCheckpointY);
		computeMypodVariables(x, y, nextCheckpointDist, nextCheckpointAngle);
		computeOpponentVariables(opponentX, opponentY);
		
		fprintf(stderr, " my pod orientation is %d degree\n", myPod.orientation);
		fprintf(stderr, " my pod distance is %d\n", myPod.distance);
		fprintf(stderr, " my pod direction is %d\n", myPod.direction);
		fprintf(stderr, " my pod speed is %d\n", myPod.speed);
		fprintf(stderr, " my pod booster is %d\n", myPod.booster);
		fprintf(stderr, " lap index is %d\n", lapIndex);
		fprintf(stderr, " checkpoint index is %d\n", cpIndex);
		
		/*----control algorithm------------------------------------------------------------------------------------------*/
		
		//followOpponent(); //for testing purpose only
		followCheckpoint();
		
		/*----sending outputs--------------------------------------------------------------------------------------------*/
		/* Giving output parameters to the game
           Write an action using printf(). DON'T FORGET THE TRAILING \n
           To debug: fprintf(stderr, "Debug messages...\n");
        */
        
		/*
           You have to output the target position
           followed by the power (0 <= thrust <= 100) or "BOOST"
           i.e.: "x y thrust"
        */
        printf("%d %d %s\n", myPod.goToXY.x, myPod.goToXY.y, myPod.thrustPower);
        //printf("%d %d %s\n", nextCheckpointX, nextCheckpointY, "100");

		// Update backtrack variables for my pod
		myPod.prevLoc = myPod.currLoc;
		opPod.prevLoc = opPod.currLoc;
		// Update backtract variables for opponent
    }

    return 0;
}

/*-------------------------MAIN FUNCTION PROTOTYPE---------------------------------------------------------------------*/
void initialize(void)
{
    // game variables
    nCheckpoints = 0;
    lapIndex = 0;
    lapFlag = 0;

	// my pod variables
    myPod.prevLoc.x = -1;
    myPod.prevLoc.y = -1;
    myPod.booster = 1;

	// opponent variables
}

void computeMapVariables(int checkpointX, int checkpointY)
// Input = current checkpoint coordinate   
// Output = lapIndex, cpIndex, nCheckpoints 
{
	//check if this checkpoint is included in database
    cpIndex = recognizedCheckpoint(checkpointX, checkpointY);
    
    //checkpoint not found - add it in database
    if (cpIndex == -1) //checkpoint has not register yet
    {
        addCheckpoint(checkpointX, checkpointY); //update nCheckpoints also
        cpIndex = nCheckpoints - 1;
    }    
    //checkpoint found - check if it is the start of new lap
    else
    {
    	//currently at first checkpoint == start of new lap
    	if (cpIndex == nCheckpoints - 1 && lapFlag == 0)
    	{
    		lapIndex++;
    		lapFlag = 1;
		}
		if (cpIndex < nCheckpoints - 1 )
		{
			lapFlag = 0;
		}
	}
}

void computeMypodVariables(int xo, int yo, int dist, int angle)
// Input = current location (xo, yo),  distance to checkpoint (dist), 
//			angle between the orientation and displacement to checkpoint (angle) in degree 
// Output = Velocity, Orientation, Distance, Direction
{		
	myPod.currLoc.x = xo;
	myPod.currLoc.y = yo;
	
    if (myPod.prevLoc.x == -1 || myPod.prevLoc.y == -1)
    {
		myPod.prevLoc = myPod.currLoc;
    }	
    myPod.velocity = createVector(myPod.prevLoc, myPod.currLoc);
	myPod.speed = vectorLength(myPod.velocity);
	myPod.direction = calcAngleBetweenVectors(myPod.velocity, createVector(myPod.currLoc, listCheckpoint[cpIndex]));
    myPod.orientation = angle;
    myPod.distance = dist;
}

void computeOpponentVariables(int xo, int yo)
// Input = current opponent coordinate
// Output = oppCurrLoc, oppCurrVelocity, oppCurrAcceleration, oppCurrDistance, oppAvailBoost
{	
	opPod.currLoc.x = xo;
	opPod.currLoc.y = yo;
	
    if (opPod.prevLoc.x == -1 || opPod.prevLoc.y == -1)
    {
		opPod.prevLoc = opPod.currLoc;
    }		
	
    opPod.velocity = createVector(opPod.prevLoc, opPod.currLoc);
	opPod.speed = vectorLength(opPod.velocity);	
}

void followOpponent(void)
{	
}

void followCheckpoint(void) // to test how to control my pod velocity, my pod distance to checkpoint and my pod orientation
{	
	if (requireBoostingSpeed() && myPod.booster == 1)
	{
		myPodGoTo(listCheckpoint[cpIndex], BOOST);
		myPod.booster = 0;			
		return;
	}
	
	if (lapIndex == 3 && cpIndex == nCheckpoints - 1)
	{		
		myPodGoTo(listCheckpoint[cpIndex], FULL_THROTTLE);
		return;
	}
	
	if (myPod.distance > CHECKPOINT_RADIUS + INBOUND_RADIUS)
	{
		if(abs(myPod.direction) > APPROACH_ANGLE && myPod.speed > MAX_APPROACH_SPEED)
		{
			myPodGoTo(listCheckpoint[cpIndex], BRAKE);	
			return;
		}
		if (abs(myPod.orientation) > MAX_ORIENTATION)
		{
			myPodGoTo(listCheckpoint[cpIndex], BRAKE);			
		}
		else
		{
			myPodGoTo(listCheckpoint[cpIndex], FULL_THROTTLE);
		}		
		return;
	}
	
	if (myPod.distance < CHECKPOINT_RADIUS + BUFFER_RADIUS && possibleForCornerTurn())
	{
		int nextCheckpoint = cpIndex + 1;
		if (nextCheckpoint == nCheckpoints)
		{
			nextCheckpoint = 0;
		}
		myPodGoTo(listCheckpoint[nextCheckpoint], FULL_THROTTLE);
		return;
	}	
	
	optimizeApproachingSpeed();
	myPodGoTo(listCheckpoint[cpIndex], myPod.thrustPower);	
}

/*-------------------------SUPPORT FUNCTION PROTOTYPE------------------------------------------------------------------*/
void myPodGoTo(planeXY destination, char * thrust)
{
	myPod.goToXY = destination;
	strcpy(myPod.thrustPower, thrust);
}

void addCheckpoint(int xo, int yo)
{	
	listCheckpoint[nCheckpoints].x = xo;
	listCheckpoint[nCheckpoints].y = yo;
	nCheckpoints++;
}

char recognizedCheckpoint(int xo, int yo)
{
	int i;
	for(i = 0; i < nCheckpoints; i++)
	{
		if( listCheckpoint[i].x == xo && listCheckpoint[i].y == yo)
		{
			return i;
		}
	}
	return -1;
}

double toRad(double angleInDegree)
{
	return angleInDegree * PI / 180;
}

int vectorLength(vector u)
{
	return trunc(sqrt(1.0 * u.dx * u.dx + 1.0 * u.dy * u.dy));
}

double dotProduct(vector u, vector v)
{
	return 1.0 * u.dx * v.dx + 1.0 * u.dy * v.dy;
}

vector subtractVectors(vector u, vector v)
{
	vector result;
	
	result.dx = u.dx - v.dx;
	result.dy = u.dy - v.dy;
	
	return result;
}

vector addVectors(vector u, vector v)
{
	vector result;
	
	result.dx = u.dx + v.dx;
	result.dy = u.dy + v.dy;
	
	return result;
}

void toString(char str[6], int n)
{
	int len, dummy;
    int i, digit;
     	
 	len = 0;
	dummy = n;    
    do
    {
        len++;
        dummy /= 10;
    }
    while(dummy != 0);
    
    for (i = 0; i < len; i++)
    {
        digit = n % 10;
        n = n / 10;
        str[len - (i + 1)] = digit + '0';
    }
    str[len] = 0;
}

vector createVector(planeXY pU, planeXY pV)
{
	vector result;
	result.dx = pV.x - pU.x;
	result.dy = pV.y - pU.y;
	return result;
}

char possibleForCornerTurn()
{
	if (abs(myPod.direction) > APPROACH_ANGLE) return 0;
	if (myPod.speed < MIN_APPROACH_SPEED) return 0;
	return 1;	
}

char requireBoostingSpeed()
{
	if (lapIndex == 3 && cpIndex == nCheckpoints - 1 && abs(myPod.orientation) < 10) return 1;
	if (abs(myPod.direction) > 10) return 0;	
	if (myPod.distance > LONG_DISTANCE && myPod.speed < MAX_APPROACH_SPEED) return 1;
	
	return 0;
}

int calcAngleBetweenVectors(vector u, vector v)
{
	double cosTheta = 1.0 * dotProduct(u, v) / (1.0 * vectorLength(u) * vectorLength(v));
	if (cosTheta >= -1 && cosTheta <= 1)
	{
		return toDegree(acos(cosTheta));
	}
	
	return 0;
}

int toDegree(double angleInRad)
{
	return trunc(angleInRad * 180 / PI);
}

void optimizeApproachingSpeed()
{
	if (myPod.direction > APPROACH_ANGLE)
	{
		strcpy(myPod.thrustPower, FULL_THROTTLE);
		return;
	}
	
	if (myPod.speed < MIN_APPROACH_SPEED)
	{
		strcpy(myPod.thrustPower, FULL_THROTTLE);
		return;
	}
	
	if (myPod.speed > MAX_APPROACH_SPEED)
	{
		strcpy(myPod.thrustPower, BRAKE);
		return;		
	}
	
	strcpy(myPod.thrustPower, HALF_THROTTLE);
}