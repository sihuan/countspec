#!/bin/bash

folders=(400.perlbench 401.bzip2 403.gcc 410.bwaves 416.gamess 429.mcf 433.milc 434.zeusmp 435.gromacs 436.cactusADM 437.leslie3d 444.namd 445.gobmk 447.dealII 450.soplex 453.povray 454.calculix 456.hmmer 458.sjeng 459.GemsFDTD 462.libquantum 464.h264ref 465.tonto 470.lbm 471.omnetpp 473.astar 481.wrf 482.sphinx3 483.xalancbmk)

for folder in "${folders[@]}"; do
    run_dir="$folder/run"
    
    if [ -d "$run_dir" ]; then
        cd "$run_dir" || exit

        for sub_dir in */; do
            cd "$sub_dir" || exit

            /home/sihuan/PLCT/spec/cpu2017/bin/specinvoke -n speccmds.cmd > mygocmd

            cd ..
        done
    else
        echo "Folder $run_dir not found."
    fi
    cd ../..
done


for folder in "${folders[@]}"; do
    run_dir="$folder/run"
    
    if [ -d "$run_dir" ]; then
        cd "$run_dir" || exit

        for sub_dir in */; do
	    mv "$sub_dir" .. || exit
        done
	echo "$run_dir done"
    else
        echo "Folder $run_dir not found."
    fi
    cd ../..
done


for folder in "${folders[@]}"; do
    run_dir="$folder/run"
    
    if [ -d "$run_dir" ]; then
	rm -r "$run_dir"
	echo "$run_dir done"
    else
        echo "Folder $run_dir not found."
    fi
done

