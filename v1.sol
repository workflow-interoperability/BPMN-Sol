pragma solidity ^0.4.24;

contract bpmn_v1 {

    bool[2] public TaskActivated;  //status variables
    bool public Terminated = false;

    uint[] public Task1Previous;
    uint[] public Task2Previous;
    uint[] public Task1Next;
    uint[] public Task2Next;
    
    event Task1Finished(uint time);
    event Task2Finished(uint time);
    event ProcessFinished(uint time);

    function Init() {
        TaskActivated[0] = true;   
        TaskActivated[1] = false;    
    }
    
    function Task1() {  //first task
        require(TaskActivated[0] == true);  //conformance checking

        //add next & previous elements, needed by all tasks except the last one
        Task1Next.push(2);  //add next elements
        Task2Previous.push(1);  //add next elements' previous elements

        //activate next elements, needed by all tasks except the last one
        for (uint i = 0; i < Task1Next.length; i++) {  
            TaskActivated[Task1Next[i] - 1] = true;  
        }

        Task1Finished(now);  //send event
    }

    function Task2() {  //last task  
        require(TaskActivated[1] == true);  //conformance checking

       //deactivate previous elements, needed by all tasks except the first one
        for (uint i = 0; i < Task2Previous.length; i++) {  
            TaskActivated[Task2Previous[i] - 1] = false;  
        }

        Task2Finished(now);  //send event

        Terminated = true;  //set Terminated, needed only by the last task
    }

    function Terminate() {
        require(Terminated == true);  //conformance checking
        
        ProcessFinished(now);  //send event
    }

}
    
   