import React, {
  useEffect,
  useRef
} from 'react';
import * as d3 from 'd3'; // TODO: IMPORT ONLY REQUIRED
import { GraphExpression } from '../../types/stateExpression';
import { Point } from '../../types';
import styles from './styles.module.scss';
import './graph.css';

export interface Props {
  xScaleDomain: [number, number];
  yScaleDomain: [number, number];
  className?: string;
  graphExpression?: GraphExpression | null;
  points?: Point[];
  responsive?: boolean;
}

export const Graph: React.FC<Props> = React.memo((
  {
    points = [],
    graphExpression = {id: 1, name: '', points: []},
    className,
    xScaleDomain,
    yScaleDomain,
    responsive = false,
  }
): JSX.Element => {
  const chartRef = useRef<SVGSVGElement>(null);
  const expression = graphExpression || {id: 1, name: '', points: []};
  const height = responsive ? 500 : 100;


  useEffect(() => {
    const margin = {top: 20, right: 30, bottom: 30, left: 50};
    // const height = responsive ? 500 : parseInt(d3.select('#func-chart').style('height')) - margin.top - margin.bottom;
    const width = responsive ? parseInt(d3.select('#func-chart').style('width')) - margin.left - margin.right : 200;

    let svg: any;
    if (responsive) {
      svg = d3.select(chartRef.current)
        .attr('viewBox', `0 0 ${width + margin.left + margin.right} ${height + margin.top + margin.bottom}`)
        .append('g')
        .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')');
    } else {
      svg = d3.select(chartRef.current)
        .attr('width', width + margin.left + margin.right)
        .attr('height', height + margin.top + margin.bottom)
        .append('g')
        .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')');
    }

    //add x axis
    const xScale = d3.scaleLinear()
      .domain(xScaleDomain)
      .range([0, width]);

    //set y scale
    const yScale = d3.scaleLinear()
      .domain(yScaleDomain)
      .range([height, 0]);

    let line = d3.line()
      .x(d => xScale(d[0]))
      .y(d => yScale(d[1])); //.curve(d3.curveBasis)

    const lineGroup = svg.append('g') // add line to svg. use path-->to know svg path // must add css class line, you can try to remove it and see the result

    const tooltip = d3.select('body')
      .append('div')
      .attr('class', 'tooltip')
      .style('visibility', 'hidden')
      .text('');

    const chartPoints = lineGroup.append('g')
      .selectAll('point')
      .data(points)
      .enter()
      .append('circle')
      .attr('r', 3)
      // @ts-ignore
      .attr('cx', d => xScale(d[0]))
      // @ts-ignore
      .attr('cy', d => yScale(d[1]))
      .on('mouseover', () => tooltip.style('visibility', 'visible'))
      // @ts-ignore
      .on('mousemove', (e, d) => {
        tooltip.style('top', (e.pageY - 10) + 'px').style('left', (e.pageX + 15) + 'px');
        tooltip.text(`x: ${d[0]}, y: ${d[1]}`);
      })
      .on('mouseout', () => tooltip.style('visibility', 'hidden'));

    const funcLine = lineGroup.append('path')
      .datum(expression.points)
      .attr('fill', 'none')
      .attr('stroke', 'teal')
      .attr('stroke-width', 2)
      .attr('d', line);

    //add x and y axis
    const yAxis = d3.axisLeft(yScale).tickSize(-width);
    const yAxisGroup = svg.append('g').call(yAxis);

    const xAxis = d3.axisBottom(xScale).tickSize(-height);/*.tickFormat("");remove tick label*/
    const xAxisGroup = svg.append('g').call(xAxis).attr('transform', 'translate(0,' + height + ')');

    // add label for x axis and y axis
    svg.append('text').text('Y Label')
      .attr('x', 0 - height / 2)
      .attr('y', 0 - margin.left)
      .attr('dy', '1em')
      .style('text-anchor', 'middle')
      .attr('transform', 'rotate(-90)');
    svg.append('text').text('X Label')
      .attr('x', width / 2)
      .attr('y', height + margin.bottom)
      .style('text-anchor', 'middle');

    // add zoom
    const zoom = d3.zoom()
      .scaleExtent([0.2, 30])// less than 1 means can resize smaller than  original size
      .on('zoom', zoomed);

    function zoomed({transform}: any) {
      const new_x = transform.rescaleX(xScale);
      const new_y = transform.rescaleY(yScale);
      // update axes
      xAxisGroup.call(xAxis.scale(new_x));
      yAxisGroup.call(yAxis.scale(new_y));

      chartPoints.data(points)
        // @ts-ignore
        .attr('cx', d => new_x(d[0]))
        // @ts-ignore
        .attr('cy', d => new_y(d[1]));

      funcLine.attr('transform', transform);
    }

    const zoomer = d3.select('svg').call(zoom as any);

    //reset button
    d3.select('#reset').on('click', () => zoomer.call(zoom.transform as any, d3.zoomIdentity));

    // add clip path to the svg
    d3.select('svg').append('defs').append('clipPath').attr('id', 'clip')
      .append('rect').attr('width', width).attr('height', height);
    lineGroup.attr('clip-path', 'url(#clip)');//line group is in a fixed position and the path will be moved

    return () => {
      svg.remove();
      tooltip.remove();
    };
  }, [expression.points, points, xScaleDomain, yScaleDomain, height, responsive]);

  return (
    <div id="func-chart" className={`${styles.funcChart} ${className && className} ${!responsive && styles.funcFixedSize}`}>
      <svg ref={chartRef}/>
      <button id="reset" className={styles.resetButton}>Reset pos</button>
    </div>
  );
}, () => false);

//TODO: USE MEMO
//TODO: ALLOW DATA TO HAVE NAME (LABELS)